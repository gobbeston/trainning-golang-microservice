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
        "/v1/roles": {
            "post": {
                "description": "Create a role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "roles"
                ],
                "summary": "Create Role",
                "operationId": "post-roles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "for request tracking",
                        "name": "X-Correlation-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "All params related to role",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateRoleRequestJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateRoleResponseSwagger"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.errorResponse"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "post": {
                "description": "Create a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create User",
                "operationId": "post-users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "for request tracking",
                        "name": "X-Correlation-ID",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "All params related to user",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequestJSON"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateRoleRequestJSON": {
            "type": "object",
            "required": [
                "code",
                "displayName"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "example": "BOOM"
                },
                "displayName": {
                    "type": "string",
                    "example": "BoomDisplay"
                }
            }
        },
        "models.CreateRoleResponseJSON": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "BOOM"
                },
                "createAt": {
                    "type": "string",
                    "example": "2019-02-14T02:35:31.2296459Z"
                },
                "displayName": {
                    "type": "string",
                    "example": "BoomDisplay"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.CreateRoleResponseSwagger": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CreateRoleResponseJSON"
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "models.CreateUserRequestJSON": {
            "type": "object",
            "required": [
                "createdBy",
                "email",
                "firstName",
                "lastName",
                "password",
                "provider",
                "roleId",
                "roleTypeId",
                "statusId"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "9999/99 Parkred"
                },
                "age": {
                    "type": "string",
                    "example": "20"
                },
                "birthDate": {
                    "type": "string",
                    "example": "1999/09/09"
                },
                "createdBy": {
                    "type": "string",
                    "example": "N API"
                },
                "email": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "firstName": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "lastName": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "password": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "+66999999999"
                },
                "provider": {
                    "type": "string",
                    "example": "OWN"
                },
                "roleId": {
                    "type": "integer",
                    "example": 1
                },
                "roleTypeId": {
                    "type": "integer",
                    "example": 1
                },
                "statusId": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.CreateUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CreateUserResponseJSON"
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "models.CreateUserResponseJSON": {
            "type": "object",
            "required": [
                "createdBy",
                "email",
                "firstName",
                "lastName",
                "password",
                "provider",
                "roleId",
                "roleTypeId",
                "statusId"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "92/7 Parkred"
                },
                "age": {
                    "type": "string",
                    "example": "20"
                },
                "birthDate": {
                    "type": "string",
                    "example": "1999/09/09"
                },
                "createdBy": {
                    "type": "string",
                    "example": "N API"
                },
                "email": {
                    "type": "string",
                    "example": "tech@mail.co.th"
                },
                "firstName": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "lastName": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "password": {
                    "type": "string",
                    "example": "tech@mail.co"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "+66999999999"
                },
                "provider": {
                    "type": "string",
                    "example": "OWN"
                },
                "roleId": {
                    "type": "integer",
                    "example": 1
                },
                "roleTypeId": {
                    "type": "integer",
                    "example": 1
                },
                "statusId": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "utils.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Error message will be show here"
                },
                "status": {
                    "type": "string",
                    "example": "fail"
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
	Version:     "0.1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Trainning Golang MicroService",
	Description: "Trainning Golang MicroService.",
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
