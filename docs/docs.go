// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "leachiM2k",
            "url": "https://github.com/leachim2k/go-shorten",
            "email": "leachiM2k@leachiM2k.de"
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
        "/shorten/": {
            "post": {
                "description": "Create a new short. Create random code if not specified.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new short",
                "operationId": "create",
                "parameters": [
                    {
                        "description": "Create Request",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dataservice.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dataservice.Entity"
                        }
                    }
                }
            }
        },
        "/shorten/handle/{code}": {
            "get": {
                "description": "Return the right link for short code or \"not found\" if expired, not started or max count was reached",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "Handle a short",
                "operationId": "handle",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Link to follow",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "fail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shorten/{code}": {
            "get": {
                "description": "Get all stored information for a specified short",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get short's info",
                "operationId": "read",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dataservice.Entity"
                        }
                    },
                    "500": {
                        "description": "fail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates several fields of a short, while maintaining count, owner and creation date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a short",
                "operationId": "update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Request",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dataservice.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dataservice.Entity"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a short",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a short",
                "operationId": "delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "dataservice.CreateRequest": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expiresAt": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "maxCount": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                }
            }
        },
        "dataservice.Entity": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "code": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expiresAt": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "maxCount": {
                    "type": "integer"
                },
                "owner": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dataservice.UpdateRequest": {
            "type": "object",
            "properties": {
                "attributes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "description": {
                    "type": "string"
                },
                "expiresAt": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "maxCount": {
                    "type": "integer"
                },
                "startTime": {
                    "type": "string"
                }
            }
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Host:        "localhost:8080",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Go Shorten API",
	Description: "URL Shortener",
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
	swag.Register(swag.Name, &s{})
}
