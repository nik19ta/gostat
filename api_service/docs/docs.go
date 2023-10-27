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
            "name": "Nikita Khvatov",
            "url": "https://khvat.pro",
            "email": "nik19ta.me@gmail.com"
        },
        "license": {
            "name": "GNU Affero General Public License v3.0",
            "url": "https://github.com/nikkhvat/gostat/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/apps/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
                        "description": "Example: {\\\"successfully\\\": true, \\\"app\\\": \\\"deleted_app_id\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAppResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\": true, \\\"detail\\\": \\\"detailed error message\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAppResp"
                        }
                    }
                }
            }
        },
        "/apps/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete an application based on the application ID and user ID from the bearer token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apps"
                ],
                "summary": "Delete an application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"successfully\\\": true, \\\"app\\\": \\\"app_id\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAppResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\": true, \\\"detail\\\": \\\"detailed error message\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAppResp"
                        }
                    }
                }
            }
        },
        "/auth/confirm": {
            "post": {
                "description": "Uses the secret provided in the URL to confirm the email of an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Confirm the email of an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Secret key for account confirmation",
                        "name": "secret",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: {\\\"successful\\\":true}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthConfirmResponse"
                        }
                    },
                    "401": {
                        "description": "Example: {\\\"error\\\":\\\"Unexpected error, failed to verify account\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/confirm/mail": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Send an email with a code in order to confirm the account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Send confirmation email",
                "responses": {
                    "200": {
                        "description": "Example: {\\\"successful\\\":true}",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthConfirmResponse"
                        }
                    },
                    "400": {
                        "description": "Example: {\\\"error\\\":\\\"Unexpected error, failed to send mail\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    },
                    "401": {
                        "description": "Example: {\\\"error\\\":\\\"Invalid token\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/info": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get detailed information about a user's account and their associated applications",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Retrieve user account information",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved user information",
                        "schema": {
                            "$ref": "#/definitions/service.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters or error retrieving user information",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    },
                    "401": {
                        "description": "Example: {\\\"Invalid token\\\"}",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Uses (login or email) and password for authentication to get access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Authenticate a user and get access and refresh tokens",
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
        "/auth/password/request": {
            "post": {
                "description": "Sends a password reset email to the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Request Password Reset",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset email sent successfully",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthConfirmResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or an error occurred while sending the email",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorAuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/password/reset": {
            "post": {
                "description": "Resets the user's password using a secret token sent to their email.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ResetConfirmPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset successfully",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessAuthResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or an error occurred while resetting the password",
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
        },
        "/stats/get/visits": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
        "/stats/set/visit": {
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
        "/stats/set/visit/extend": {
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
        "http.ErrorAppResp": {
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
        "http.SuccessAppResponse": {
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
        "http.SuccessAuthConfirmResponse": {
            "type": "object",
            "properties": {
                "successful": {
                    "description": "Indicates whether the confirmation was successful\nexample: true",
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
        },
        "service.ResetConfirmPasswordRequest": {
            "type": "object",
            "properties": {
                "mail": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "service.ResetPasswordRequest": {
            "type": "object",
            "properties": {
                "mail": {
                    "description": "The email address associated with the account\nrequired: true\nexample: user@example.com",
                    "type": "string"
                }
            }
        },
        "service.UserApp": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Timestamp of application creation\nexample: \"2023-10-22T01:47:40+04:00\"",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier of the application\nexample: \"8d8da463-767a-488c-9cc6-45dc35346507\"",
                    "type": "string"
                },
                "image": {
                    "description": "Image or icon of the application\nexample: \"default\"",
                    "type": "string"
                },
                "name": {
                    "description": "Name of the application\nexample: \"nikkhvat\"",
                    "type": "string"
                },
                "url": {
                    "description": "URL of the application\nexample: \"https://nik.khvat.pro\"",
                    "type": "string"
                }
            }
        },
        "service.UserInfo": {
            "type": "object",
            "properties": {
                "account_confirmed": {
                    "description": "Account confirmation status\nexample: false",
                    "type": "boolean"
                },
                "apps": {
                    "description": "List of applications associated with the user",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.UserApp"
                    }
                },
                "avatar": {
                    "description": "Avatar URL\nexample: \"\"",
                    "type": "string"
                },
                "created_at": {
                    "description": "Timestamp of account creation\nexample: \"2023-10-22 00:49:35\"",
                    "type": "string"
                },
                "email": {
                    "description": "Email address of the user\nexample: \"nik19ta.me1@gmail.com\"",
                    "type": "string"
                },
                "first_name": {
                    "description": "First name of the user\nexample: \"Nikita\"",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier of the user\nexample: 168",
                    "type": "integer"
                },
                "last_name": {
                    "description": "Last name of the user\nexample: \"Khvatov\"",
                    "type": "string"
                },
                "login": {
                    "description": "Login name of the user\nexample: \"nik19ta.me1\"",
                    "type": "string"
                },
                "middle_name": {
                    "description": "Middle name of the user\nexample: \"Dmitrievich\"",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Gostat",
	Description:      "Statistics Service - gostat. A microservice-based service, written in Golang and TypeScript.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
