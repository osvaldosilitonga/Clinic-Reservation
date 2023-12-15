// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Zi.Care",
            "url": "www.zicare.com",
            "email": "zicare@mail.com"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/employee/login": {
            "post": {
                "description": "Login for employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Employee Login",
                "parameters": [
                    {
                        "description": "Employee Credentials",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SwLoginRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    }
                }
            }
        },
        "/employee/register": {
            "post": {
                "description": "Register new employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Employee Register",
                "parameters": [
                    {
                        "description": "Employee Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.SwUserRegisterRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login for patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Patient Login",
                "parameters": [
                    {
                        "description": "Patient Credentials",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SwLoginRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "Logout the currently authenticated user and clears the authorization cookie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Logout the user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SwLoginRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register new patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Default"
                ],
                "summary": "Patient Register",
                "parameters": [
                    {
                        "description": "Patient Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.SwUserRegisterRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrWebResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.EmployeeLoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                }
            }
        },
        "dto.EmployeeRegisterReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                }
            }
        },
        "dto.ErrWebResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "detail": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.SwLoginRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.SwUserRegisterRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.UserResponse"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.UserLoginReq": {
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
                    "maxLength": 32,
                    "minLength": 6
                }
            }
        },
        "dto.UserRegisterReq": {
            "type": "object",
            "required": [
                "address",
                "birth",
                "email",
                "gender",
                "name",
                "password",
                "phone"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "maxLength": 128,
                    "minLength": 8
                },
                "birth": {
                    "type": "string",
                    "maxLength": 10,
                    "minLength": 10
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 4
                },
                "name": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 6
                },
                "phone": {
                    "type": "string",
                    "maxLength": 13,
                    "minLength": 10
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "BETA",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Zi.Care Clinic - API Documentation",
	Description:      "Application for managing clinic data and appointments",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
