// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/loki/auth/login/": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "responses": {
                    "200": {
                        "description": "{}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/auth/register/": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register",
                "responses": {
                    "201": {
                        "description": "{}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/auth/userinfo/": {
            "get": {
                "description": "Get userinfo by jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get userinfo by jwt token",
                "responses": {
                    "200": {
                        "description": "{}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/context/": {
            "get": {
                "description": "limit 2000",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get loki log context from grafana loki and accepts the following query parameters in the URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The start time for the query as a nanosecond Unix epoch",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The end time for the query as a nanosecond Unix epoch",
                        "name": "end",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/export/": {
            "get": {
                "description": "file log (max count 50000)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Download loki log to log file and accepts the following query parameters in the URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The start time for the query as a nanosecond Unix epoch",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The end time for the query as a nanosecond Unix epoch",
                        "name": "end",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The filter condition",
                        "name": "filter",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The log level",
                        "name": "level",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The order to all results",
                        "name": "dsc",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/history": {
            "get": {
                "description": "Get loki query history labels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get loki query history labels",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Every page count",
                        "name": "page_size",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page index",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/history/create": {
            "post": {
                "description": "Create loki query history api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create loki query history api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The query label value dict",
                        "name": "label_json",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The query filter value list",
                        "name": "filter_json",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/history/delete/:id": {
            "delete": {
                "description": "Delete loki query history labels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete loki query history labels",
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/label/values/": {
            "get": {
                "description": "limit 2000",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves the list of known values for a given label within a given time span. It accepts the following query parameters in the URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The start time for the query as a nanosecond Unix epoch",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The end time for the query as a nanosecond Unix epoch",
                        "name": "end",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The label value",
                        "name": "label",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/labels/": {
            "get": {
                "description": "limit 2000",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves the list of known values for a given label within a given time span. It accepts the following query parameters in the URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The start time for the query as a nanosecond Unix epoch",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The end time for the query as a nanosecond Unix epoch",
                        "name": "end",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/list/": {
            "get": {
                "description": "default limit 2000",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Used to do a query over a range of time and accepts the following query parameters in the URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The start time for the query as a nanosecond Unix epoch",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The end time for the query as a nanosecond Unix epoch",
                        "name": "end",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The new query to all results",
                        "name": "all",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The order to all results",
                        "name": "dsc",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The filter condition",
                        "name": "filter",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The log level",
                        "name": "level",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The max number of entries to return",
                        "name": "limit",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/settings/load/": {
            "get": {
                "description": "Load Settings for UI",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "LoadSettings",
                "responses": {
                    "200": {
                        "description": "{}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/snapshot": {
            "get": {
                "description": "Get loki query result snapshot list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get loki query result snapshot list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Every page count",
                        "name": "page_size",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page index",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/snapshot/create": {
            "post": {
                "description": "Create loki query result snapshot result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create loki query result snapshot result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Snapshot filename",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Snapshot result temp file",
                        "name": "tmp_file",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Snapshot query result start time",
                        "name": "start_time",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Snapshot query result end time",
                        "name": "end_time",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/snapshot/delete/:id": {
            "delete": {
                "description": "Delete loki query result snapshot file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete loki query result snapshot file",
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/loki/snapshot/detail/:id": {
            "get": {
                "description": "Get loki query result snapshot detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get loki query result snapshot detail",
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ws/tail/": {
            "get": {
                "description": "limit 2000",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "WebSocket endpoint that will stream log messages based on a query. It accepts the following query parameters in the URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The end time for the query as a nanosecond Unix epoch",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The pod filter condition to perform",
                        "name": "pod",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The filter condition",
                        "name": "filter",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The log level",
                        "name": "level",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "The max number of entries to return",
                        "name": "limit",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "[]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "dagger backend api",
	Description: "this is dagger backend api server",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
