package utils

import (
	"dagger/backend/databases"
	"dagger/backend/models"
	"dagger/backend/runtime"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	conf "github.com/prometheus/alertmanager/config"
	"go.uber.org/zap"
)

var (
	SMTPHost       string
	SMTPPort       string
	SMTPAuthUser   string
	SMTPAuthPass   string
	SMTPFrom       string
	RepeatInterval string
	ResolveTimeout string
)

func init() {
	SMTPHost, _ = runtime.Cfg.GetValue("alertmanager", "smtp_smarthost")
	SMTPPort, _ = runtime.Cfg.GetValue("alertmanager", "smtp_smartport")
	SMTPAuthUser, _ = runtime.Cfg.GetValue("alertmanager", "smtp_auth_username")
	SMTPAuthPass, _ = runtime.Cfg.GetValue("alertmanager", "smtp_auth_password")
	SMTPFrom, _ = runtime.Cfg.GetValue("alertmanager", "smtp_from")
}

func FlushConf2Alertmanager(content string) error {
	err := ioutil.WriteFile("conf/alertmanager.yml", []byte(content), 0644)
	return err
}

func ReloadAlertmanager() error {
	address, _ := runtime.Cfg.GetValue("alertmanager", "address")

	url := fmt.Sprintf("/-/reload")
	var err error

	_, err = HttpRequest(fmt.Sprintf("%s%s", address, url), "POST", nil, nil, "json")
	if err != nil {
		Log4Zap(zap.WarnLevel).Warn(fmt.Sprintf("reload alertmanager error %s", err))
		return err
	}

	return nil
}

func LoadAlertmanagerConf() ([]byte, error) {
	address, _ := runtime.Cfg.GetValue("alertmanager", "address")

	url := fmt.Sprintf("/api/v2/status")

	data, err := HttpRequest(fmt.Sprintf("%s%s", address, url), "GET", nil, nil, "json")
	if err != nil {
		Log4Zap(zap.WarnLevel).Warn(fmt.Sprintf("load alertmanager conf error %s", err))
		return nil, err
	}

	var jsonRes map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonRes)
	if err != nil {
		Log4Zap(zap.WarnLevel).Warn(fmt.Sprintf("Unmarshal alermanager conf response error %s", err))
		return nil, err
	}

	b, _ := json.Marshal(jsonRes["config"].(map[string]interface{})["original"])

	return b, nil
}

func Push2Alertmanager(data string) error {
	address, _ := runtime.Cfg.GetValue("alertmanager", "address")

	url := fmt.Sprintf("/api/v2/alerts")

	data, err := HttpRequest(fmt.Sprintf("%s%s", address, url), "POST", nil, data, "json")
	if err != nil {
		Log4Zap(zap.WarnLevel).Warn(fmt.Sprintf("push to alertmanager error %s", err))
		return err
	}

	return nil
}

func DynamicAlertmanagerConf() error {
	ruleMapI, _ := databases.GC.Get("rule-map")
	if ruleMapI != nil {

		data, err := LoadAlertmanagerConf()
		if err != nil {
			return err
		}

		var y interface{}
		err = yaml.Unmarshal(data, &y)
		if err != nil {
			return err
		}

		cf, err := conf.Load(fmt.Sprintf("%s", y))
		if err != nil {
			return err
		}

		cf.Global = &conf.GlobalConfig{
			SMTPSmarthost: conf.HostPort{
				Host: SMTPHost,
				Port: SMTPPort,
			},
			SMTPFrom:         SMTPFrom,
			SMTPAuthUsername: SMTPAuthUser,
			SMTPAuthPassword: conf.Secret(SMTPAuthPass),
		}
		receivers := []*conf.Receiver{}
		route := conf.Route{
			Routes: []*conf.Route{},
		}

		allowSignUp, _ := runtime.Cfg.Bool("users", "allow_sign_up")

		ruleMap := ruleMapI.(map[string]models.LogRule)
		for _, rule := range ruleMap {

			userstr := ""
			mails := []string{}
			for _, group := range rule.Groups {
				for _, user := range group.LogUserGroup.Users {
					if strings.Index(userstr, user.User.Email) == -1 {
						mails = append(mails, user.User.Email)
						userstr += fmt.Sprintf("%s,", user.User.Email)
					}
				}
			}
			if strings.Index(userstr, rule.User.Email) == -1 {
				mails = append(mails, rule.User.Email)
			}

			to := ""
			if allowSignUp {
				to = strings.Join(mails, ",")
			} else {
				to, _ = runtime.Cfg.GetValue("global", "to")
			}

			receive := conf.Receiver{
				Name: rule.Name,
				EmailConfigs: []*conf.EmailConfig{
					{
						To: to,
					},
				},
			}
			receivers = append(receivers, &receive)
			route.Routes = append(route.Routes, &conf.Route{
				Receiver: rule.Name,
				Match:    map[string]string{"alertname": rule.Name},
			})
		}

		cf.Receivers = receivers
		cf.Route = &route
		content := cf.String()

		err = FlushConf2Alertmanager(content)
		if err != nil {
			return err
		}

		err = ReloadAlertmanager()
		if err != nil {
			return err
		}
	}

	return nil
}
