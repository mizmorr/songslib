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
        "/song": {
            "put": {
                "description": "Update an existing song record with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Update song information",
                "parameters": [
                    {
                        "description": "Song data to be updated",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SongRequestUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid input or request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new song record in the database with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Create a new song",
                "parameters": [
                    {
                        "description": "Song data to be created",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SongRequestCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created song",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid input or request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a song by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Delete a song",
                "parameters": [
                    {
                        "description": "Song to delete",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SongRequestDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request, invalid data",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        },
        "/song/pages": {
            "get": {
                "description": "Retrieve a list of songs based on filtering criteria and pagination options.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get filtered and paginated list of songs",
                "parameters": [
                    {
                        "type": "string",
                        "name": "band",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "lyrics",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved filtered songs list",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid input or request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        },
        "/song/verses": {
            "get": {
                "description": "Retrieve all verses of a song based on the song details and pagination options provided.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "songs"
                ],
                "summary": "Get verses of a song",
                "parameters": [
                    {
                        "type": "string",
                        "name": "band",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "lyrics",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of song verses",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Verse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid input or request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "details": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "model.SongRequestCreate": {
            "type": "object",
            "required": [
                "band",
                "name"
            ],
            "properties": {
                "band": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.SongRequestDelete": {
            "type": "object",
            "properties": {
                "band": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lyrics": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.SongRequestUpdate": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "band": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lyrics": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Verse": {
            "type": "object",
            "properties": {
                "band": {
                    "type": "string"
                },
                "lines": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "number": {
                    "type": "integer"
                },
                "song": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a simple restful service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
