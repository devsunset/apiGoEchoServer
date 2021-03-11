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
        "contact": {
            "name": "API Support",
            "url": "https://devsunset.github.io",
            "email": "devsunset@gmail.com"
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
        "/getLocation": {
            "get": {
                "description": "get tesla Chargers Location.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Tesla App"
                ],
                "summary": "Tesla Chargers Location.",
                "operationId": "getLocation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "asia_pacific",
                        "name": "region",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "South Korea",
                        "name": "country",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "서울특별시",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "supercharger",
                        "name": "location_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "28865",
                        "name": "nid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Locations"
                        }
                    }
                }
            }
        },
        "/getLocationCity": {
            "get": {
                "description": "get tesla Chargers Location City.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Tesla App"
                ],
                "summary": "Tesla Chargers Location City.",
                "operationId": "getLocationCity",
                "parameters": [
                    {
                        "type": "string",
                        "description": "South Korea",
                        "name": "country",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Citys"
                        }
                    }
                }
            }
        },
        "/getLocationCountry": {
            "get": {
                "description": "get tesla Chargers Location Country.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Tesla App"
                ],
                "summary": "Tesla Chargers Location Country.",
                "operationId": "getLocationCountry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "asia_pacific",
                        "name": "region",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Countrys"
                        }
                    }
                }
            }
        },
        "/getLocationRegion": {
            "get": {
                "description": "get tesla Chargers Location Region.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Tesla App"
                ],
                "summary": "Tesla Chargers Location Region.",
                "operationId": "getLocationRegion",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Regions"
                        }
                    }
                }
            }
        },
        "/getLocationType": {
            "get": {
                "description": "get tesla Chargers Location Type.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Tesla App"
                ],
                "summary": "Tesla Chargers Location Type.",
                "operationId": "getLocationType",
                "parameters": [
                    {
                        "type": "string",
                        "description": "South Korea",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "서울특별시",
                        "name": "city",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Location_types"
                        }
                    }
                }
            }
        },
        "/getPathParameters/{id}": {
            "get": {
                "description": "get value from path",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "example get value from path",
                "operationId": "getPathParameters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/getQueryParameters": {
            "get": {
                "description": "get value from the query string",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "example get value from the query string",
                "operationId": "getQueryParameters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "2",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/healthCheck": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API Tesla App"
                ],
                "summary": "Show the status of server.",
                "operationId": "HealthCheck",
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
        "/postSave": {
            "post": {
                "description": "get value from the form",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "example get value from the form",
                "operationId": "postSave",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "2",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/postSavefile": {
            "post": {
                "description": "get value/file from the form",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "example get value/file from the form",
                "operationId": "postSavefile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "2",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "-",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "get value from the form/json",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "example get value from the form/json",
                "operationId": "users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "2",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.City": {
            "type": "object",
            "properties": {
                "City": {
                    "type": "string"
                }
            }
        },
        "main.Citys": {
            "type": "object",
            "properties": {
                "Citys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.City"
                    }
                }
            }
        },
        "main.Country": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                }
            }
        },
        "main.Countrys": {
            "type": "object",
            "properties": {
                "countrys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Country"
                    }
                }
            }
        },
        "main.Emails_Data": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                }
            }
        },
        "main.Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "address_line_1": {
                    "type": "string"
                },
                "address_line_2": {
                    "type": "string"
                },
                "address_notes": {
                    "type": "string"
                },
                "amenities": {
                    "type": "string"
                },
                "baidu_lat": {
                    "type": "string"
                },
                "baidu_lng": {
                    "type": "string"
                },
                "chargers": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "common_name": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "destination_charger_logo": {
                    "type": "string"
                },
                "destination_website": {
                    "type": "string"
                },
                "directions_link": {
                    "type": "string"
                },
                "emails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Emails_Data"
                    }
                },
                "geocode": {
                    "type": "string"
                },
                "hours": {
                    "type": "string"
                },
                "is_gallery": {
                    "type": "string"
                },
                "kiosk_pin_x": {
                    "type": "string"
                },
                "kiosk_pin_y": {
                    "type": "string"
                },
                "kiosk_zoom_pin_x": {
                    "type": "string"
                },
                "kiosk_zoom_pin_y": {
                    "type": "string"
                },
                "latitude": {
                    "type": "string"
                },
                "location_id": {
                    "type": "string"
                },
                "location_type": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "longitude": {
                    "type": "string"
                },
                "nid": {
                    "type": "string"
                },
                "open_soon": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "province_state": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "sales_phone": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Sales_Phone_Data"
                    }
                },
                "sales_representative": {
                    "type": "string"
                },
                "sl_translate": {
                    "type": "string"
                },
                "sub_region": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "trt_id": {
                    "type": "string"
                }
            }
        },
        "main.Location_type": {
            "type": "object",
            "properties": {
                "Location_type": {
                    "type": "string"
                }
            }
        },
        "main.Location_types": {
            "type": "object",
            "properties": {
                "Location_types": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Location_type"
                    }
                }
            }
        },
        "main.Locations": {
            "type": "object",
            "properties": {
                "locations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Location"
                    }
                }
            }
        },
        "main.Region": {
            "type": "object",
            "properties": {
                "region": {
                    "type": "string"
                }
            }
        },
        "main.Regions": {
            "type": "object",
            "properties": {
                "regions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Region"
                    }
                }
            }
        },
        "main.Sales_Phone_Data": {
            "type": "object",
            "properties": {
                "label": {
                    "type": "string"
                },
                "line_below": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                }
            }
        }
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
	Host:        "193.123.252.22:8282",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "apiServer  API",
	Description: "This is a apiServer.",
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
