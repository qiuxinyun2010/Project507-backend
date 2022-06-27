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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/menu/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取后台管理菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/article": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取文章列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page_num",
                        "name": "page_num",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "uid",
                        "name": "uid",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "标签",
                        "name": "tag_name",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "image",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/article/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/article/{id}/addTags": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加文章标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "标签ID",
                        "name": "tag_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/article/{id}/deleteTags": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除文章标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "标签ID",
                        "name": "tag_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/article/{id}/like": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加文章点赞",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/article/{id}/recover": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "恢复文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/tag": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取标签列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "UserId",
                        "name": "uid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "10001": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "10006": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        },
        "/api/v1/tag/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取该标签的所有文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Modifiedby",
                        "name": "modified_by",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "10007": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "10008": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        },
        "/api/v1/tag/{id}/clear": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "清理标签(硬删除)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/tag/{id}/recover": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "恢复标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/{id}/follow": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "关注用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "关注用户ID",
                        "name": "follow_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "类型",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/upload": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "上传图片",
                "parameters": [
                    {
                        "type": "file",
                        "description": "image",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "20005": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "20006": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "20007": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改用户密码",
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
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "20009": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新用户",
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
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        },
        "/user/{id}/refreshToken": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "更新Token",
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
                        "description": "uuid",
                        "name": "uuid",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_http.ResponseJSON"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin_http.ResponseJSON": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
