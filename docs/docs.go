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
        "contact": {
            "name": "Yusuf BULUT",
            "url": "https://www.linkedin.com/in/yusufbu1ut/",
            "email": "yusufblt10@outlook.com"
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
        "/basket": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket"
                ],
                "summary": "Gets all basket items for the authed customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/basket.ResponseBasket"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket"
                ],
                "summary": "Adds the given items to basket",
                "parameters": [
                    {
                        "description": "Takes the products and adds them",
                        "name": "basket-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/basket.ResponseProduct"
                            }
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket"
                ],
                "summary": "Deletes given basket products id, it can be with query or request body",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Takes the basket product infos and deletes it, product id is needed",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "description": "Takes the basket product infos and deletes it, product id is needed",
                        "name": "product-request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/basket.ResponseProduct"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket"
                ],
                "summary": "Updates given basket item, only amount can be updated",
                "parameters": [
                    {
                        "description": "Takes the basket infos and updates it, basket id and to update amount is needed",
                        "name": "basket-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/basket.ResponseProduct"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/basket/complete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket"
                ],
                "summary": "Completes baskets to orders with basket products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/category": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Gets all categories with pagination parameters page and size",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Index",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pagination.Pages"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create categories with given csv file uploads it after that reads and creates (adding form-data not implemented for swagger)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login/admin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Admins login with email and password",
                "parameters": [
                    {
                        "description": "Login process takes admin' email and password. Checks the inputs in database and returns JWT token.",
                        "name": "login-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.RequestLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login/customer": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Customers login with email and password",
                "parameters": [
                    {
                        "description": "Login process takes customer' email and password. Checks the inputs in database and returns JWT token.",
                        "name": "login-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.RequestLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Gets all order items for the authed customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.ResponseOrder"
                        }
                    }
                }
            }
        },
        "/order/cancel": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Cancels the orders with given order code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Takes the order codes and cancel them",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Gets all products with pagination parameters page and size",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Index",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pagination.Pages"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Updates given product",
                "parameters": [
                    {
                        "description": "Takes the product infos and updates it. Product id, name, category id and code needed",
                        "name": "product-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ResponseProduct"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Creates products with the given request",
                "parameters": [
                    {
                        "description": "Takes the products and adds them to db",
                        "name": "product-request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ResponseProduct"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Deletes given product id, it can be with query or request body",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Takes the product id and deletes it, product id is needed",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "description": "Takes the product infos and deletes it, product id is needed",
                        "name": "product-request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/product.ResponseProduct"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/product/search": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Gets all products with search parameters name, amount or categoryId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "CategoryId",
                        "name": "categoryId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Amount",
                        "name": "amount",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.ResponseProduct"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SignUp"
                ],
                "summary": "Customers can sign up with username,email and password(Needed fields)",
                "parameters": [
                    {
                        "description": "Sing-up takes  username, email and password necessarily. Other fields not necessary but customers can add them too. Checks the customer in database adds it and returns JWT token.",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sign_up.RequestCustomer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "basket.ResponseBasket": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/basket.ResponseProduct"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "basket.ResponseProduct": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "login.RequestLogin": {
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
        "order.ResponseOrder": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.ResponseProduct"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "order.ResponseProduct": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "order-time": {
                    "type": "string"
                }
            }
        },
        "pagination.Pages": {
            "type": "object",
            "properties": {
                "items": {},
                "page": {
                    "type": "integer"
                },
                "pageCount": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "totalCount": {
                    "type": "integer"
                }
            }
        },
        "product.ResponseProduct": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "category-id": {
                    "type": "integer"
                },
                "code": {
                    "type": "integer"
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
                "price": {
                    "type": "number"
                }
            }
        },
        "sign_up.RequestCustomer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Golang Bootcamp Ending Project- Basket API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
