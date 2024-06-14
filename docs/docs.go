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
            "email": "fiber@swagger.io"
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
        "/insert": {
            "post": {
                "description": "Input data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Insert data presensi.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPresensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/presensi": {
            "get": {
                "description": "Mengambil semua data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get All Data Presensi.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                }
            }
        },
        "/presensi/{id}": {
            "get": {
                "description": "Ambil per ID data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get By ID Data Presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.JamKerja": {
            "type": "object",
            "properties": {
                "durasi": {
                    "type": "integer",
                    "example": 8
                },
                "gmt": {
                    "type": "integer",
                    "example": 7
                },
                "hari": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jam_keluar": {
                    "type": "string",
                    "example": "16:00"
                },
                "jam_masuk": {
                    "type": "string",
                    "example": "08:00"
                },
                "piket_tim": {
                    "type": "string",
                    "example": "Piket Z"
                },
                "shift": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "controller.Karyawan": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "123456789"
                },
                "hari_kerja": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jabatan": {
                    "type": "string",
                    "example": "Anonymous"
                },
                "jam_kerja": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.JamKerja"
                    }
                },
                "nama": {
                    "type": "string",
                    "example": "Tes Swagger"
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.Presensi": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "123456789"
                },
                "biodata": {
                    "$ref": "#/definitions/controller.Karyawan"
                },
                "checkin": {
                    "type": "string",
                    "example": "Masuk"
                },
                "datetime": {
                    "type": "string",
                    "format": "date-time",
                    "example": "2024-09-01T00:00:00Z"
                },
                "latitude": {
                    "type": "number",
                    "example": 123.12
                },
                "location": {
                    "type": "string",
                    "example": "Bandung"
                },
                "longitude": {
                    "type": "number",
                    "example": 123.11
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.ReqKaryawan": {
            "type": "object",
            "properties": {
                "hari_kerja": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jabatan": {
                    "type": "string",
                    "example": "Anonymous"
                },
                "jam_kerja": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.JamKerja"
                    }
                },
                "nama": {
                    "type": "string",
                    "example": "Tes Swagger"
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.ReqPresensi": {
            "type": "object",
            "properties": {
                "biodata": {
                    "$ref": "#/definitions/controller.ReqKaryawan"
                },
                "checkin": {
                    "type": "string",
                    "example": "Masuk"
                },
                "latitude": {
                    "type": "number",
                    "example": 123.12
                },
                "location": {
                    "type": "string",
                    "example": "Bandung"
                },
                "longitude": {
                    "type": "number",
                    "example": 123.11
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ws-bargana2024-abecd99cf108.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
