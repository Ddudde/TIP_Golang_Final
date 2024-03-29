// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/auth/login": {
            "post": {
                "description": "Authorization in system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authorization",
                "parameters": [
                    {
                        "description": "Login and pass",
                        "name": "main",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.LoginResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registration new account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Registration",
                "parameters": [
                    {
                        "description": "Login and pass",
                        "name": "main",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.RegisterResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Отправка каталога магазина",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get Product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.GetProductResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.GetProductResponse"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получение профиля юзера",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Get Profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.GetProfileResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/practs_mainRoute_pb.GetProfileResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "practs_mainRoute_pb.GetProductResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/practs_mainRoute_pb.Prodct"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "practs_mainRoute_pb.GetProfileResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "logname": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "practs_mainRoute_pb.LoginRequest": {
            "type": "object",
            "properties": {
                "logname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "practs_mainRoute_pb.LoginResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "practs_mainRoute_pb.Prodct": {
            "type": "object",
            "properties": {
                "gram": {
                    "type": "integer"
                },
                "img": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "zag": {
                    "type": "string"
                }
            }
        },
        "practs_mainRoute_pb.RegisterRequest": {
            "type": "object",
            "properties": {
                "logname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "practs_mainRoute_pb.RegisterResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"HTTP"},
	Title:            "MainRouteService",
	Description:      "Route in HTTP to Services(AuthM, ProductM, ProfileM).\\n\\nMainRoute: doc.json\\nAuthService: http://localhost:3000/docs1/mainRoute/pb/auth.swagger.json\\nProductService: http://localhost:3000/docs1/mainRoute/pb/product.swagger.json\\nProfileService: http://localhost:3000/docs1/mainRoute/pb/profile.swagger.json",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
