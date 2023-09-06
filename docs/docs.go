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
        "/{instanceId}/chat/messages": {
            "post": {
                "description": "Returns chat messages from the specified WhatsApp instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Chat"
                ],
                "summary": "Get WhatsApp Chat Messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of chat messages",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/{instanceId}/chat/send/audio": {
            "post": {
                "description": "Sends an audio message on WhatsApp using the specified instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Chat"
                ],
                "summary": "Send Audio Message on WhatsApp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Audio message body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.audioMessageBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message Send Response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/{instanceId}/chat/send/image": {
            "post": {
                "description": "Sends an image message on WhatsApp using the specified instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Chat"
                ],
                "summary": "Send Image Message on WhatsApp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Image message body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.imageMessageBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message Send Response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/{instanceId}/chat/send/text": {
            "post": {
                "description": "Sends a text message on WhatsApp using the specified instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Chat"
                ],
                "summary": "Send Text Message on WhatsApp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Text message body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.textMessageBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message Send Response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{instanceId}/check/phones": {
            "post": {
                "description": "Verifies if the phone numbers in the provided list are registered WhatsApp users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Phone Verification"
                ],
                "summary": "Check Phones on WhatsApp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Phone list",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.phoneCheckBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of verified numbers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/{instanceId}/contact/info": {
            "get": {
                "description": "Retrieves contact information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Contact"
                ],
                "summary": "Get Contact Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Phone",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Contact Information",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/{instanceId}/logout": {
            "post": {
                "description": "Logs out from the specified WhatsApp instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Logout"
                ],
                "summary": "Logout from WhatsApp",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Logout successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/{instanceId}/profile": {
            "get": {
                "description": "Retrieves profile information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Profile"
                ],
                "summary": "Get Profile Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile Information",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/{instanceId}/qrcode": {
            "get": {
                "description": "Returns a QR code to initiate WhatsApp login.",
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "WhatsApp Login"
                ],
                "summary": "Get WhatsApp QR Code",
                "responses": {
                    "200": {
                        "description": "PNG image containing the QR code",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/{instanceId}/status": {
            "get": {
                "description": "Returns the status of the specified WhatsApp instance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "WhatsApp Status"
                ],
                "summary": "Get WhatsApp Instance Status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Instance ID",
                        "name": "instanceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status Response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.audioMessageBody": {
            "type": "object",
            "properties": {
                "base64": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "controllers.imageMessageBody": {
            "type": "object",
            "properties": {
                "base64": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "controllers.phoneCheckBody": {
            "type": "object",
            "properties": {
                "phones": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controllers.textMessageBody": {
            "type": "object",
            "properties": {
                "phone": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8900",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "ZapMeow API",
	Description:      "API to handle multiple WhatsApp instances",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
