// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "A method for obtaining information about a user using his passport data. Defined in the task conditions.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Get user info",
                "operationId": "ger-user-info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "passportSerie",
                        "name": "passportSerie",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "passportNumber",
                        "name": "passportNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/entity.People"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/info/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Filtering is possible by the fields \"surname\", \"name\", \"patronymic\", \"address\". The field(s) for sorting and its direction are specified.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Receiving data about all users with filtering and sorting.",
                "operationId": "get-all-user-info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filtering by fields: surname, name, patronymic, address",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sorting by fields: surname, name, patronymic, address",
                        "name": "sortProperty",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sorting direction DESC and ASC",
                        "name": "sortDirection",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "output limit - maximum value is 10.",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.People"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/info/effort": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Obtaining labor costs by user for a period/ task is the sum of hours and minutes, sorted from the highest cost to the least",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Obtaining labor costs for a user in a certain period.",
                "operationId": "get-user-effort",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "beginning of period in format 2024-07-03",
                        "name": "beginning",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end of period in format 2024-07-04",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/v1.getUserEffort.responce"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/people/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "The format is specified in the task conditions. Creates a new user using passport data. Returns the ID of the created user. To further fill/change data, the Update method works.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Create person",
                "operationId": "create-person",
                "parameters": [
                    {
                        "description": "passportNumber",
                        "name": "passportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.userInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/people/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Based on accepted passport data, removes the user from the system.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Delete information about a person",
                "operationId": "delete-person",
                "parameters": [
                    {
                        "description": "passportNumber",
                        "name": "passportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.userInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/people/update": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates the user data to the specified one. user verification is carried out using the passport data in the body of the request. If successful, returns a message and user ID.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "summary": "Updates information about a person",
                "operationId": "update-person",
                "parameters": [
                    {
                        "description": "passportNumber",
                        "name": "passportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.userInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "surname user",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name user",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "patronymic user",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "address user",
                        "name": "address",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "We receive all the tasks and set a limit on page output.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get all tasks",
                "operationId": "get-all-task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit on page outpu",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new task using the parameters and returns its ID on success. By default, a new task has an internal status of \"planned\"",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create task",
                "operationId": "create-task",
                "parameters": [
                    {
                        "description": "You need to give the task a name, its importance (high or low(default)) and a description.",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Deleting a task using the ID obtained from the parameter.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Deleting a task",
                "operationId": "delete-task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task ID to delete",
                        "name": "taskID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "We get the task by its ID in the request parameters.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Retrieving a task by its ID",
                "operationId": "get-task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Yes",
                        "name": "taskId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tracker/start": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Starts counting down the task completion time for the user. Inside, the task status changes from \"planned\" to \"accepted\". If successful, we receive the tracker ID.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tracker"
                ],
                "summary": "Starts timing on a task for the user.",
                "operationId": "start-tracker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "specify task ID",
                        "name": "task_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "specify user ID",
                        "name": "people_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tracker/stop": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Ends the task timer for the user. Inside, the task status changes from “accepted” to “completed”.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tracker"
                ],
                "summary": "Stops the execution time of a task for the user.",
                "operationId": "stop-tracker",
                "parameters": [
                    {
                        "type": "string",
                        "description": "specify task ID",
                        "name": "task_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "specify user ID",
                        "name": "people_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignIn",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SignUp",
                "operationId": "create_account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Manager"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/v1.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Effort": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "task_id": {
                    "type": "string"
                },
                "total_time": {
                    "type": "string"
                }
            }
        },
        "entity.Manager": {
            "type": "object",
            "required": [
                "managername",
                "name",
                "password"
            ],
            "properties": {
                "managername": {
                    "type": "string",
                    "example": "Manager"
                },
                "name": {
                    "type": "string",
                    "example": "Andrew"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty"
                }
            }
        },
        "entity.People": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "г. Москва, ул. Ленина, д. 5, кв. 1"
                },
                "name": {
                    "type": "string",
                    "example": "Иван"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Иванович"
                },
                "surname": {
                    "type": "string",
                    "example": "Иванов"
                }
            }
        },
        "entity.Task": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "A very important task"
                },
                "importance": {
                    "type": "string",
                    "example": "low or high"
                },
                "name": {
                    "type": "string",
                    "example": "T-001"
                }
            }
        },
        "v1.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.getUserEffort.responce": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "effort": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Effort"
                    }
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.signInInput": {
            "type": "object",
            "required": [
                "managername",
                "password"
            ],
            "properties": {
                "managername": {
                    "type": "string",
                    "example": "Manager"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty"
                }
            }
        },
        "v1.userInput": {
            "type": "object",
            "required": [
                "passportNumber"
            ],
            "properties": {
                "passportNumber": {
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
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Effective Mobile API",
	Description:      "API Server for test work",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
