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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "健康接口",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/ticket/user": {
            "post": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User management"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "{` + "`" + `userId` + "`" + `: xxx, ` + "`" + `password` + "`" + `: xxx}",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/grabbing_ticket_service.UserType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User management"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "{` + "`" + `userId` + "`" + `: xxx}",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/grabbing_ticket_service.DeleteUserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/ticket/user-list": {
            "get": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User management"
                ],
                "summary": "用户列表",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/web/auth/check-token": {
            "get": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Web用户"
                ],
                "summary": "检查token有效性",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token=XXX",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/web/login": {
            "post": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Web用户"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "{` + "`" + `name` + "`" + `: xxx, ` + "`" + `password` + "`" + `: xxx}",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/web/register": {
            "post": {
                "description": "返回ok",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Web用户"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "{` + "`" + `name` + "`" + `: xxx, ` + "`" + `password` + "`" + `: xxx}",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RegisterInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.RegisterInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "grabbing_ticket_service.DeleteUserParam": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "string"
                }
            }
        },
        "grabbing_ticket_service.UserType": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "model.LoginReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "yannscrapy API Docs",
	Description: "This is yannscrapy.",
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
