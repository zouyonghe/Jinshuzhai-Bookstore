// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/zouyonghe",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/zouyonghe",
            "email": "1259085392z@gmail.com"
        },
        "license": {
            "name": "GPLv3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/admin/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a user account specified by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/admin"
                ],
                "summary": "Update a user account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the ID of the specified user to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user information include username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":5}}",
                        "schema": {
                            "$ref": "#/definitions/user.UpdateResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a user by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/admin"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the ID of the specified user to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":5}}",
                        "schema": {
                            "$ref": "#/definitions/user.DeleteResponse"
                        }
                    }
                }
            }
        },
        "/user/common": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update the current user information by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/common"
                ],
                "summary": "Update the current user information",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SelfUpdRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":6,\"username\":\"夏秀兰\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.SelfUpdResponse"
                        }
                    }
                }
            }
        },
        "/user/common/": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "SelfDel deletes the user of token specified",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user/common"
                ],
                "summary": "SelfDel deletes the user of token specified",
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":8}}",
                        "schema": {
                            "$ref": "#/definitions/user.SelfDelResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login a user account with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login  a user account",
                "parameters": [
                    {
                        "description": "Login account",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":\"7\",\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NTA0MzkzNDAsImlkIjo3LCJuYmYiOjE2NTA0MzkzNDAsInJvbGUiOiJnZW5lcmFsIiwidXNlcm5hbWUiOiLpob7no4oifQ.ZqeFEugcvTS2Rgq0qR4Na49-rkye6CoXPV_R9ub-QYQ\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Create a new user by username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "user information include username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"userId\":\"7\",\"username\":\"顾磊\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.CreateRequest": {
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
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.DeleteResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        },
        "user.LoginRequest": {
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
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "user.SelfDelResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        },
        "user.SelfUpdRequest": {
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
        "user.SelfUpdResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UpdateRequest": {
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
        "user.UpdateResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
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
	Version:          "0.3",
	Host:             "127.0.0.1:8081",
	BasePath:         "/v1",
	Schemes:          []string{"https"},
	Title:            "Jinshuzhai-Bookstore",
	Description:      "The jinshuzhai bookstore api server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
