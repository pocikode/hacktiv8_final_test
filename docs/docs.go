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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Agus Supriyatna",
            "email": "supriyatnaagus@outlook.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comments": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Comment",
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Comment",
                "parameters": [
                    {
                        "description": "Comment",
                        "name": "Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqComment"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/comments/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Comment",
                "parameters": [
                    {
                        "description": "Comment",
                        "name": "Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqComment"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/photos": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Photo",
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Photo",
                "parameters": [
                    {
                        "description": "Photo",
                        "name": "Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPhoto"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/photos/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Photo",
                "parameters": [
                    {
                        "description": "Photo",
                        "name": "Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqPhoto"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photo"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Photo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/socialmedias": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socmed"
                ],
                "summary": "Socmed",
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socmed"
                ],
                "summary": "Socmed",
                "parameters": [
                    {
                        "description": "Photo",
                        "name": "Socmed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqSocmed"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/socialmedias/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socmed"
                ],
                "summary": "Socmed",
                "parameters": [
                    {
                        "description": "Photo",
                        "name": "Socmed",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqSocmed"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Socmed ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socmed"
                ],
                "summary": "Socmed",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Socmed ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "description": "Update User",
                        "name": "Update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqUserUpdate"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete User",
                "responses": {}
            }
        },
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqLogin"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "Register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqRegister"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.ReqComment": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                }
            }
        },
        "models.ReqLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.ReqPhoto": {
            "type": "object",
            "required": [
                "photo_url",
                "title"
            ],
            "properties": {
                "caption": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.ReqRegister": {
            "type": "object",
            "required": [
                "age",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "minimum": 9
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ReqSocmed": {
            "type": "object",
            "required": [
                "name",
                "social_media_url"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "models.ReqUserUpdate": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
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
	Version:          "",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "",
	Description:      "MyGram - Hacktiv8 Final Test",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}