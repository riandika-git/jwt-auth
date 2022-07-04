// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customer/detail": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get Customer Detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Customer Detail",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CustomerDetailDTO"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "Login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDetail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.CustomerLoginDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/token/refresh": {
            "post": {
                "description": "Refresh Access Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Refresh Access Token",
                "parameters": [
                    {
                        "description": "Refresh Token",
                        "name": "RefreshToken",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserLoginDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/voucher-group": {
            "get": {
                "description": "Voucher Group List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Voucher Group List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.VoucherGroupDTO"
                            }
                        }
                    }
                }
            }
        },
        "/voucher-purchase": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Purchase Voucher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Purchase Voucher",
                "parameters": [
                    {
                        "description": "Voucher Group ID",
                        "name": "Finance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PurchaseVoucherRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.VoucherPurchaseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CustomerDetailDTO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CustomerLoginDTO": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "accessUuid": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDetail": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.PurchaseVoucherRequestDTO": {
            "type": "object",
            "properties": {
                "voucherGroupId": {
                    "type": "integer"
                }
            }
        },
        "dto.RefreshToken": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "dto.UserLoginDTO": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "accessUuid": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "deposit": {
                    "type": "number"
                },
                "email": {
                    "type": "string"
                },
                "firstLogin": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                },
                "whatsapp": {
                    "type": "string"
                }
            }
        },
        "dto.VoucherGroupDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "voucherGroupName": {
                    "type": "string"
                }
            }
        },
        "dto.VoucherPurchaseDTO": {
            "type": "object",
            "properties": {
                "customerId": {
                    "type": "integer"
                },
                "customerName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "purchaseDate": {
                    "type": "string"
                },
                "voucherCode": {
                    "type": "string"
                },
                "voucherGroupId": {
                    "type": "integer"
                },
                "voucherGroupName": {
                    "type": "string"
                }
            }
        },
        "helper.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger API Notification",
	Description: "This is a microservice jwt-auth.",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
