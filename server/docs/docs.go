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
        "/ping": {
            "get": {
                "description": "gin server의 헬스를 체크합니다.",
                "produces": [
                    "application/json"
                ],
                "summary": "Server Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Pong"
                        }
                    }
                }
            }
        },
        "/users/signup": {
            "post": {
                "description": "사용자를 생성합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "회원가입",
                "parameters": [
                    {
                        "description": "회원가입을 위한 정보",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/users/1"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "여기에 에러 메시지가 전달됩니다."
                }
            }
        },
        "models.Pong": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "pong"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "passwordConfirm"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user01@test.com"
                },
                "name": {
                    "type": "string",
                    "example": "사용자01"
                },
                "password": {
                    "type": "string",
                    "example": "test1!"
                },
                "passwordConfirm": {
                    "type": "string",
                    "example": "test1!"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
