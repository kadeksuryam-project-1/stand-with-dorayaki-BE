// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kadek Surya Mahardika",
            "email": "kadeksuryam@gmail.com"
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
        "/v1/dorayakis": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorayakis"
                ],
                "summary": "Get Dorayakis",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetDorayakisResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorayakis"
                ],
                "summary": "Create Dorayaki",
                "parameters": [
                    {
                        "type": "string",
                        "description": "flavor",
                        "name": "flavor",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CreateDorayakiResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/v1/dorayakis/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorayakis"
                ],
                "summary": "Get Dorayaki",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetDorayakiResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorayakis"
                ],
                "summary": "Update Dorayaki",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "flavor",
                        "name": "flavor",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.UpdateDorayakiResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorayakis"
                ],
                "summary": "Delete Dorayaki",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.DeleteDorayakiResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/v1/stocks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Get stocks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetStocksResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Update Stock",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "dorayaki_id",
                        "name": "dorayaki_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "store_id",
                        "name": "store_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "stock",
                        "name": "stock",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.StockRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.UpdateStockResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/v1/stores": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stores"
                ],
                "summary": "Get stores",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetStoresResponseDTO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stores"
                ],
                "summary": "Create store",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "street",
                        "name": "street",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "subdistrict",
                        "name": "subdistrict",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "district",
                        "name": "district",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "province",
                        "name": "province",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CreateStoreResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/v1/stores/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stores"
                ],
                "summary": "Get store",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetStoreResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stores"
                ],
                "summary": "Update store",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "street",
                        "name": "street",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "subdistrict",
                        "name": "subdistrict",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "district",
                        "name": "district",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "province",
                        "name": "province",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.UpdateStoreResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stores"
                ],
                "summary": "Delete store",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.DeleteStoreResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.CreateDorayakiResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.Dorayaki"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.CreateStoreResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.Store"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.DeleteDorayakiResponseDTO": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.DeleteStoreResponseDTO": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.Dorayaki": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "flavor": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "schema.DorayakiStoreStock": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "dorayaki": {
                    "$ref": "#/definitions/schema.Dorayaki"
                },
                "dorayaki_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "store": {
                    "$ref": "#/definitions/schema.Store"
                },
                "store_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "schema.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.GetDorayakiResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.Dorayaki"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.GetDorayakisResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Dorayaki"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.GetStocksResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.DorayakiStoreStock"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.GetStoreResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.Store"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.GetStoresResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Store"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.StockRequestDTO": {
            "type": "object",
            "properties": {
                "stock": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "schema.Store": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "subdistrict": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "schema.UpdateDorayakiResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.Dorayaki"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.UpdateStockResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.DorayakiStoreStock"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schema.UpdateStoreResponseDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schema.Store"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Stand with Dorayaki API",
	Description:      "Stand with Dorayaki API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
