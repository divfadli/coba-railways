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
        "/age-rating-categories": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get a list of AgeRatingCategory.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get all AgeRatingCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AgeRatingCategory"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Creating a new AgeRatingCategory.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Create New AgeRatingCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "the body to create a new AgeRatingCategory",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ageRatingCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            }
        },
        "/age-rating-categories/{id}": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get an AgeRatingCategory by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get AgeRatingCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AgeRatingCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a AgeRatingCategory by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Delete one AgeRatingCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AgeRatingCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update AgeRatingCategory by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Update AgeRatingCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AgeRatingCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update age rating category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ageRatingCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AgeRatingCategory"
                        }
                    }
                }
            }
        },
        "/age-rating-categories/{id}/movies": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all Movies by AgeRatingCategoryId.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AgeRatingCategory"
                ],
                "summary": "Get Movies.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AgeRatingCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logging in to get jwt token to access admin or user api by roles.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login as as user.",
                "parameters": [
                    {
                        "description": "the body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/movie/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a movie by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Delete one movie.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get a list of Movies.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get all movies.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Creating a new Movie.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Create New Movie.",
                "parameters": [
                    {
                        "description": "the body to create a new movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.movieInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get a Movie by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Get Movie.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update movie by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movie"
                ],
                "summary": "Update Movie.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "movie id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update an movie",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.movieInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registering a user from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a user.",
                "parameters": [
                    {
                        "description": "the body to register a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.ageRatingCategoryInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.movieInput": {
            "type": "object",
            "properties": {
                "age_rating_category_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.AgeRatingCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "ageRatingCategoryID": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
