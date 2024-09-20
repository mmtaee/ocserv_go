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
        "/api/v1/site/": {
            "get": {
                "description": "Get site configuration",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "site"
                ],
                "summary": "Get site configuration",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Site"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "post": {
                "description": "Post site configuration",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "site"
                ],
                "summary": "Post site configuration",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "site",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/site.CreateData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Site"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "patch": {
                "description": "Update site configuration",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "site"
                ],
                "summary": "Update site configuration",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request Body",
                        "name": "site",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/site.UpdateData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Site"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/api/v1/users/": {
            "post": {
                "description": "Set up an admin or superuser during site initialization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Set up an admin user",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/api/v1/users/login/": {
            "post": {
                "description": "Login admin or staff user to get token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateLoginData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/api/v1/users/password/": {
            "patch": {
                "description": "Update admin or staff user password (self change)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update Password",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/api/v1/users/staffs/": {
            "post": {
                "description": "Create staff user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create staff user",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateStaffData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "403": {
                        "description": "Admin Permission required"
                    }
                }
            }
        },
        "/api/v1/users/staffs/:id/": {
            "delete": {
                "description": "Delete staff user(by admin)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete staff",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Admin Permission required"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/api/v1/users/staffs/:id/password/": {
            "patch": {
                "description": "Update staff user password(by admin)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update staff password",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateStaffPasswordData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "403": {
                        "description": "Admin Permission required"
                    },
                    "404": {
                        "description": "User not found"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Site": {
            "type": "object",
            "properties": {
                "captcha_secret_key": {
                    "type": "string"
                },
                "captcha_site_key": {
                    "type": "string"
                },
                "default_traffic": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "site.CreateData": {
            "type": "object",
            "required": [
                "default_traffic"
            ],
            "properties": {
                "captcha_secret_key": {
                    "type": "string"
                },
                "captcha_site_key": {
                    "type": "string"
                },
                "default_traffic": {
                    "type": "number"
                }
            }
        },
        "site.UpdateData": {
            "type": "object",
            "properties": {
                "captcha_secret_key": {
                    "type": "string"
                },
                "captcha_site_key": {
                    "type": "string"
                },
                "default_traffic": {
                    "type": "number"
                }
            }
        },
        "user.CreateData": {
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
        "user.CreateLoginData": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "remember_me": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateResponse": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateStaffData": {
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
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "expire_at": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "user.UpdateData": {
            "type": "object",
            "required": [
                "current_password",
                "new_password"
            ],
            "properties": {
                "current_password": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                }
            }
        },
        "user.UpdateStaffPasswordData": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
