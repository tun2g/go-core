// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/refresh-token": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Refresh Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh Token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.TokenResDto"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/sign-in": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "loginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginReqDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.AuthResDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/sign-up": {
            "post": {
                "description": "User Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "registerReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterReqDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/auth.AuthResDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "maximum": 50,
                        "minimum": 1,
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "ASC",
                            "DESC"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "ASC",
                            "DESC"
                        ],
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order_field",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/auth.UserResDto"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/users/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get me",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.UserResDto"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/exception.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AuthResDto": {
            "type": "object",
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/auth.TokenResDto"
                },
                "user": {
                    "$ref": "#/definitions/auth.UserResDto"
                }
            }
        },
        "auth.LoginReqDto": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "auth.RegisterReqDto": {
            "type": "object",
            "required": [
                "email",
                "fullName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "auth.TokenResDto": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "auth.UserResDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "constants.Order": {
            "type": "string",
            "enum": [
                "ASC",
                "DESC"
            ],
            "x-enum-varnames": [
                "ASC",
                "DESC"
            ]
        },
        "exception.ErrorDetail": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "issue": {
                    "type": "string"
                },
                "issueId": {
                    "type": "string"
                }
            }
        },
        "exception.HttpError": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/exception.ErrorDetail"
                    }
                },
                "message": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
	Schemes:          []string{},
	Title:            "Swagger UI",
	Description:      "Golang Gin Boilerplate.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
