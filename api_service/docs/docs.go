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
        "/api/apps/create": {
            "post": {
                "description": "Create a new application with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apps"
                ],
                "summary": "Create a new application",
                "parameters": [
                    {
                        "description": "Create App payload",
                        "name": "CreateAppRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.CreateAppRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"successfully\\\": true, \\\"app\\\": \\\"new_app_id\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAppCreateResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\": true, \\\"detail\\\": \\\"detailed error message\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAppCreateResponse"
                        }
                    }
                }
            }
        },
        "/api/stats/get/visits": {
            "get": {
                "description": "Retrieves visits data for a specific application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stats"
                ],
                "summary": "Retrieve visits for an application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application ID",
                        "name": "app",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved data. The structure of the data depends on the application.",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\": true, \\\"detail\\\": \\\"Detailed error message\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorGetVisitResponse"
                        }
                    }
                }
            }
        },
        "/api/stats/set/visit": {
            "put": {
                "description": "Registers a new visit or extends an existing session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stats"
                ],
                "summary": "Set a new visit session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Unique (1 or 0)",
                        "name": "un",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UTM parameters",
                        "name": "utm",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Visited URL",
                        "name": "url",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page Title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Session ID",
                        "name": "session",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Application ID",
                        "name": "app_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"successfully\\\": true, \\\"session\\\": \\\"session_id\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessSetVisitResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\": true, \\\"detail\\\": \\\"Detailed error message\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorSetVisitResponse"
                        }
                    }
                }
            }
        },
        "/api/stats/set/visit/extend": {
            "put": {
                "description": "Extends the session for a particular visit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stats"
                ],
                "summary": "Extend a visit session",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Visit Session ID",
                        "name": "session",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"successfully\\\": true}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessVisitExtendResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\": true, \\\"detail\\\": \\\"Detailed error message\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorVisitExtendResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Uses login and password for authentication to get an access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Authenticate a user and get an access token",
                "parameters": [
                    {
                        "description": "Login payload",
                        "name": "LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Uses the refresh token to generate a new access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Refresh the access token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Refresh token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"access_token\\\":\\\"your_new_generated_token\\\", \\\"refresh_token\\\":\\\"your_refresh_token\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthResponse"
                        }
                    },
                    "401": {
                        "description": "Example: {\\\"error\\\":\\\"Invalid refresh token\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/registration": {
            "post": {
                "description": "Register a new user with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Registration payload",
                        "name": "RegistrationRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.RegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"token\\\":\\\"YOUR_GENERATED_TOKEN\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\":\\\"any error\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    },
                    "409": {
                        "description": "Example: {\\\"error\\\":\\\"User with the same email already exists\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.CreateAppRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Application name\nexample: \"My App\"",
                    "type": "string"
                },
                "url": {
                    "description": "Application URL\nexample: \"https://myapp.com\"",
                    "type": "string"
                }
            }
        },
        "http.ErrorAppCreateResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "Detailed error message\nexample: \"detailed error message\"",
                    "type": "string"
                },
                "error": {
                    "description": "Indicates an error occurred\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "http.ErrorAuthResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "Detailed error message\nin: body\nexample: \"login or password is not correct\"",
                    "type": "string"
                }
            }
        },
        "http.ErrorGetVisitResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "Detailed error message\nexample: \"Detailed error message\"",
                    "type": "string"
                },
                "error": {
                    "description": "Indicates an error occurred\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "http.ErrorSetVisitResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "Detailed error message\nexample: \"Detailed error message\"",
                    "type": "string"
                },
                "error": {
                    "description": "Indicates an error occurred\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "http.ErrorVisitExtendResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "Detailed error message\nexample: \"Detailed error message\"",
                    "type": "string"
                },
                "error": {
                    "description": "Indicates an error occurred\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "http.SuccessAppCreateResponse": {
            "type": "object",
            "properties": {
                "app": {
                    "description": "New application ID\nexample: \"new_app_id\"",
                    "type": "string"
                },
                "successfully": {
                    "description": "Indicates successful app creation\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "http.SuccessAuthResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "The generated JWT\nin: body\nexample: \"your_generated_token\"",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "The generated Refresh\nin: body\nexample: \"your_refresh_token\"",
                    "type": "string"
                }
            }
        },
        "http.SuccessSetVisitResponse": {
            "type": "object",
            "properties": {
                "session": {
                    "description": "Session ID for the visit\nexample: \"session_id\"",
                    "type": "string"
                },
                "successfully": {
                    "description": "Indicates successful visit setting or extending\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "http.SuccessVisitExtendResponse": {
            "type": "object",
            "properties": {
                "successfully": {
                    "description": "Indicates successful visit extension\nexample: true",
                    "type": "boolean"
                }
            }
        },
        "service.LoginRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "service.RegistrationRequest": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GoStat API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}