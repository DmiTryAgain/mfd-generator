// Code generated by zenrpc; DO NOT EDIT.

package api

import (
	"context"
	"encoding/json"

	"github.com/semrush/zenrpc/v2"
	"github.com/semrush/zenrpc/v2/smd"

	"github.com/vmkteam/mfd-generator/mfd"
)

var RPC = struct {
	ProjectService struct{ Open, Update, Save, Tables string }
	PublicService  struct{ GoPGVersions, Modes, SearchTypes, Types string }
	XMLService     struct{ NSMapping, GenerateEntity, LoadEntity, SaveEntity string }
	XMLLangService struct{ LoadTranslation, TranslateEntity string }
	XMLVTService   struct{ GenerateEntity, LoadEntity, SaveEntity string }
}{
	ProjectService: struct{ Open, Update, Save, Tables string }{
		Open:   "open",
		Update: "update",
		Save:   "save",
		Tables: "tables",
	},
	PublicService: struct{ GoPGVersions, Modes, SearchTypes, Types string }{
		GoPGVersions: "gopgversions",
		Modes:        "modes",
		SearchTypes:  "searchtypes",
		Types:        "types",
	},
	XMLService: struct{ NSMapping, GenerateEntity, LoadEntity, SaveEntity string }{
		NSMapping:      "nsmapping",
		GenerateEntity: "generateentity",
		LoadEntity:     "loadentity",
		SaveEntity:     "saveentity",
	},
	XMLLangService: struct{ LoadTranslation, TranslateEntity string }{
		LoadTranslation: "loadtranslation",
		TranslateEntity: "translateentity",
	},
	XMLVTService: struct{ GenerateEntity, LoadEntity, SaveEntity string }{
		GenerateEntity: "generateentity",
		LoadEntity:     "loadentity",
		SaveEntity:     "saveentity",
	},
}

func (ProjectService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"Open": {
				Description: `Loads project from file`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "filePath",
						Optional:    false,
						Description: `the path to mfd file`,
						Type:        smd.String,
					},
					{
						Name:        "connection",
						Optional:    false,
						Description: `connection string`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Project`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"name": {
							Description: ``,
							Type:        smd.String,
						},
						"namespaces": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"type": smd.String,
							},
						},
						"languages": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"type": smd.String,
							},
						},
						"goPGVer": {
							Description: ``,
							Type:        smd.Integer,
						},
						"customTypes": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.CustomTypes",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.CustomTypes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"dbType": {
									Description: ``,
									Type:        smd.String,
								},
								"goType": {
									Description: ``,
									Type:        smd.String,
								},
								"goImport": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"Update": {
				Description: `Updates project in memory`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "project",
						Optional:    false,
						Description: `Project`,
						Type:        smd.Object,
						Properties: map[string]smd.Property{
							"name": {
								Description: ``,
								Type:        smd.String,
							},
							"namespaces": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"type": smd.String,
								},
							},
							"languages": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"type": smd.String,
								},
							},
							"goPGVer": {
								Description: ``,
								Type:        smd.Integer,
							},
							"customTypes": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.CustomTypes",
								},
							},
						},
						Definitions: map[string]smd.Definition{
							"mfd.CustomTypes": {
								Type: "object",
								Properties: map[string]smd.Property{
									"dbType": {
										Description: ``,
										Type:        smd.String,
									},
									"goType": {
										Description: ``,
										Type:        smd.String,
									},
									"goImport": {
										Description: ``,
										Type:        smd.String,
									},
								},
							},
						},
					},
				},
			},
			"Save": {
				Description: `Saves project from memory to disk`,
				Parameters:  []smd.JSONSchema{},
			},
			"Tables": {
				Description: `Gets all tables from database`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of tables`,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s ProjectService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.ProjectService.Open:
		var args = struct {
			FilePath   string `json:"filePath"`
			Connection string `json:"connection"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"filePath", "connection"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Open(args.FilePath, args.Connection))

	case RPC.ProjectService.Update:
		var args = struct {
			Project mfd.Project `json:"project"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"project"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.Update(args.Project))

	case RPC.ProjectService.Save:
		resp.Set(s.Save())

	case RPC.ProjectService.Tables:
		resp.Set(s.Tables())

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (PublicService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"GoPGVersions": {
				Description: `Gets all supported go-pg versions`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of versions`,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.Integer,
					},
				},
			},
			"Modes": {
				Description: `Gets all available entity modes`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of modes`,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
			"SearchTypes": {
				Description: `Gets all available search types`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of search types`,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
			"Types": {
				Description: `Gets std types`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `list of types`,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"type": smd.String,
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s PublicService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}

	switch method {
	case RPC.PublicService.GoPGVersions:
		resp.Set(s.GoPGVersions())

	case RPC.PublicService.Modes:
		resp.Set(s.Modes())

	case RPC.PublicService.SearchTypes:
		resp.Set(s.SearchTypes())

	case RPC.PublicService.Types:
		resp.Set(s.Types())

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (XMLService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"NSMapping": {
				Description: `Saves project at filepath location`,
				Parameters:  []smd.JSONSchema{},
				Returns: smd.JSONSchema{
					Description: `NSMapping`,
					Optional:    false,
					Type:        smd.Array,
					Items: map[string]string{
						"$ref": "#/definitions/NSMapping",
					},
					Definitions: map[string]smd.Definition{
						"NSMapping": {
							Type: "object",
							Properties: map[string]smd.Property{
								"table": {
									Description: ``,
									Type:        smd.String,
								},
								"namespace": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"GenerateEntity": {
				Description: `Gets xml for selected table`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "table",
						Optional:    false,
						Description: `selected table name`,
						Type:        smd.String,
					},
					{
						Name:        "namespace",
						Optional:    false,
						Description: `namespace of the new entity`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Entity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"name": {
							Description: ``,
							Type:        smd.String,
						},
						"namespace": {
							Description: ``,
							Type:        smd.String,
						},
						"table": {
							Description: ``,
							Type:        smd.String,
						},
						"attributes": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Attributes",
							},
						},
						"searches": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Searches",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.Attributes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"dbName": {
									Description: ``,
									Type:        smd.String,
								},
								"isArray": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"disablePointer": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"dbType": {
									Description: ``,
									Type:        smd.String,
								},
								"goType": {
									Description: ``,
									Type:        smd.String,
								},
								"pk": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"fk": {
									Description: ``,
									Type:        smd.String,
								},
								"nullable": {
									Description: ``,
									Type:        smd.String,
								},
								"addable": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"updatable": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"min": {
									Description: ``,
									Type:        smd.Integer,
								},
								"max": {
									Description: ``,
									Type:        smd.Integer,
								},
								"defaultVal": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
						"mfd.Searches": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"attrName": {
									Description: ``,
									Type:        smd.String,
								},
								"searchType": {
									Description: ``,
									Type:        smd.String,
								},
								"goType": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"LoadEntity": {
				Description: `Gets xml for selected entity in project file`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Optional:    false,
						Description: `namespace of the entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Optional:    false,
						Description: `the name of the entity`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Entity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"name": {
							Description: ``,
							Type:        smd.String,
						},
						"namespace": {
							Description: ``,
							Type:        smd.String,
						},
						"table": {
							Description: ``,
							Type:        smd.String,
						},
						"attributes": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Attributes",
							},
						},
						"searches": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.Searches",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.Attributes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"dbName": {
									Description: ``,
									Type:        smd.String,
								},
								"isArray": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"disablePointer": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"dbType": {
									Description: ``,
									Type:        smd.String,
								},
								"goType": {
									Description: ``,
									Type:        smd.String,
								},
								"pk": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"fk": {
									Description: ``,
									Type:        smd.String,
								},
								"nullable": {
									Description: ``,
									Type:        smd.String,
								},
								"addable": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"updatable": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"min": {
									Description: ``,
									Type:        smd.Integer,
								},
								"max": {
									Description: ``,
									Type:        smd.Integer,
								},
								"defaultVal": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
						"mfd.Searches": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"attrName": {
									Description: ``,
									Type:        smd.String,
								},
								"searchType": {
									Description: ``,
									Type:        smd.String,
								},
								"goType": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"SaveEntity": {
				Description: `Gets xml for selected entity in project file`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "entity",
						Optional:    true,
						Description: `Entity`,
						Type:        smd.Object,
						Properties: map[string]smd.Property{
							"name": {
								Description: ``,
								Type:        smd.String,
							},
							"namespace": {
								Description: ``,
								Type:        smd.String,
							},
							"table": {
								Description: ``,
								Type:        smd.String,
							},
							"attributes": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.Attributes",
								},
							},
							"searches": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.Searches",
								},
							},
						},
						Definitions: map[string]smd.Definition{
							"mfd.Attributes": {
								Type: "object",
								Properties: map[string]smd.Property{
									"name": {
										Description: ``,
										Type:        smd.String,
									},
									"dbName": {
										Description: ``,
										Type:        smd.String,
									},
									"isArray": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"disablePointer": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"dbType": {
										Description: ``,
										Type:        smd.String,
									},
									"goType": {
										Description: ``,
										Type:        smd.String,
									},
									"pk": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"fk": {
										Description: ``,
										Type:        smd.String,
									},
									"nullable": {
										Description: ``,
										Type:        smd.String,
									},
									"addable": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"updatable": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"min": {
										Description: ``,
										Type:        smd.Integer,
									},
									"max": {
										Description: ``,
										Type:        smd.Integer,
									},
									"defaultVal": {
										Description: ``,
										Type:        smd.String,
									},
								},
							},
							"mfd.Searches": {
								Type: "object",
								Properties: map[string]smd.Property{
									"name": {
										Description: ``,
										Type:        smd.String,
									},
									"attrName": {
										Description: ``,
										Type:        smd.String,
									},
									"searchType": {
										Description: ``,
										Type:        smd.String,
									},
									"goType": {
										Description: ``,
										Type:        smd.String,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s XMLService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.XMLService.NSMapping:
		resp.Set(s.NSMapping())

	case RPC.XMLService.GenerateEntity:
		var args = struct {
			Table     string `json:"table"`
			Namespace string `json:"namespace"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"table", "namespace"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.GenerateEntity(args.Table, args.Namespace))

	case RPC.XMLService.LoadEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.LoadEntity(args.Namespace, args.Entity))

	case RPC.XMLService.SaveEntity:
		var args = struct {
			Entity *mfd.Entity `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.SaveEntity(args.Entity))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (XMLLangService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"LoadTranslation": {
				Description: `Loads full translation of project`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "language",
						Optional:    false,
						Description: `language`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `Translation`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"language": {
							Description: ``,
							Type:        smd.String,
						},
						"namespaces": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.TranslationNamespace",
							},
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.TranslationNamespace": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"entities": {
									Description: ``,
									Type:        smd.Array,
									Items: map[string]string{
										"$ref": "#/definitions/mfd.TranslationEntity",
									},
								},
							},
						},
						"mfd.TranslationEntity": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"key": {
									Description: ``,
									Type:        smd.String,
								},
								"crumbs": {
									Description: ``,
									Ref:         "#/definitions/mfd.XMLMap",
									Type:        smd.Object,
								},
								"form": {
									Description: ``,
									Ref:         "#/definitions/mfd.XMLMap",
									Type:        smd.Object,
								},
								"list": {
									Description: ``,
									Ref:         "#/definitions/mfd.TranslationList",
									Type:        smd.Object,
								},
							},
						},
					},
				},
			},
			"TranslateEntity": {
				Description: `Translates entity`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Optional:    false,
						Description: `namespace of the vt entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Optional:    false,
						Description: `vt entity from vt.xml`,
						Type:        smd.String,
					},
					{
						Name:        "language",
						Optional:    false,
						Description: `language`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `TranslationEntity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"name": {
							Description: ``,
							Type:        smd.String,
						},
						"key": {
							Description: ``,
							Type:        smd.String,
						},
						"crumbs": {
							Description: ``,
							Ref:         "#/definitions/mfd.XMLMap",
							Type:        smd.Object,
						},
						"form": {
							Description: ``,
							Ref:         "#/definitions/mfd.XMLMap",
							Type:        smd.Object,
						},
						"list": {
							Description: ``,
							Ref:         "#/definitions/mfd.TranslationList",
							Type:        smd.Object,
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.XMLMap": {
							Type:       "object",
							Properties: map[string]smd.Property{},
						},
						"mfd.TranslationList": {
							Type: "object",
							Properties: map[string]smd.Property{
								"title": {
									Description: ``,
									Type:        smd.String,
								},
								"filter": {
									Description: ``,
									Ref:         "#/definitions/mfd.XMLMap",
									Type:        smd.Object,
								},
								"headers": {
									Description: ``,
									Ref:         "#/definitions/mfd.XMLMap",
									Type:        smd.Object,
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s XMLLangService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.XMLLangService.LoadTranslation:
		var args = struct {
			Language string `json:"language"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"language"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.LoadTranslation(args.Language))

	case RPC.XMLLangService.TranslateEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
			Language  string `json:"language"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity", "language"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.TranslateEntity(args.Namespace, args.Entity, args.Language))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}

func (XMLVTService) SMD() smd.ServiceInfo {
	return smd.ServiceInfo{
		Description: ``,
		Methods: map[string]smd.Service{
			"GenerateEntity": {
				Description: `Gets xml for selected table`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Optional:    false,
						Description: `namespace of the base entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Optional:    false,
						Description: `base entity from namespace.xml`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `VTEntity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"name": {
							Description: ``,
							Type:        smd.String,
						},
						"terminalPath": {
							Description: ``,
							Type:        smd.String,
						},
						"attributes": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.VTAttributes",
							},
						},
						"template": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.TmplAttributes",
							},
						},
						"mode": {
							Description: ``,
							Type:        smd.String,
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.VTAttributes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"attrName": {
									Description: ``,
									Type:        smd.String,
								},
								"searchName": {
									Description: ``,
									Type:        smd.String,
								},
								"summary": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"search": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"max": {
									Description: ``,
									Type:        smd.Integer,
								},
								"min": {
									Description: ``,
									Type:        smd.Integer,
								},
								"required": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"validate": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
						"mfd.TmplAttributes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"vtAttrName": {
									Description: ``,
									Type:        smd.String,
								},
								"list": {
									Description: `show in list`,
									Type:        smd.Boolean,
								},
								"fkOpts": {
									Description: `how to show fks`,
									Type:        smd.String,
								},
								"form": {
									Description: `show in object editor`,
									Type:        smd.String,
								},
								"search": {
									Description: `input type in search`,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"LoadEntity": {
				Description: `Gets xml for selected entity in project file`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Optional:    false,
						Description: `namespace of the vt entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Optional:    false,
						Description: `the name of the vt entity`,
						Type:        smd.String,
					},
				},
				Returns: smd.JSONSchema{
					Description: `VTEntity`,
					Optional:    true,
					Type:        smd.Object,
					Properties: map[string]smd.Property{
						"name": {
							Description: ``,
							Type:        smd.String,
						},
						"terminalPath": {
							Description: ``,
							Type:        smd.String,
						},
						"attributes": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.VTAttributes",
							},
						},
						"template": {
							Description: ``,
							Type:        smd.Array,
							Items: map[string]string{
								"$ref": "#/definitions/mfd.TmplAttributes",
							},
						},
						"mode": {
							Description: ``,
							Type:        smd.String,
						},
					},
					Definitions: map[string]smd.Definition{
						"mfd.VTAttributes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"attrName": {
									Description: ``,
									Type:        smd.String,
								},
								"searchName": {
									Description: ``,
									Type:        smd.String,
								},
								"summary": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"search": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"max": {
									Description: ``,
									Type:        smd.Integer,
								},
								"min": {
									Description: ``,
									Type:        smd.Integer,
								},
								"required": {
									Description: ``,
									Type:        smd.Boolean,
								},
								"validate": {
									Description: ``,
									Type:        smd.String,
								},
							},
						},
						"mfd.TmplAttributes": {
							Type: "object",
							Properties: map[string]smd.Property{
								"name": {
									Description: ``,
									Type:        smd.String,
								},
								"vtAttrName": {
									Description: ``,
									Type:        smd.String,
								},
								"list": {
									Description: `show in list`,
									Type:        smd.Boolean,
								},
								"fkOpts": {
									Description: `how to show fks`,
									Type:        smd.String,
								},
								"form": {
									Description: `show in object editor`,
									Type:        smd.String,
								},
								"search": {
									Description: `input type in search`,
									Type:        smd.String,
								},
							},
						},
					},
				},
			},
			"SaveEntity": {
				Description: `Gets xml for selected entity in project file`,
				Parameters: []smd.JSONSchema{
					{
						Name:        "namespace",
						Optional:    false,
						Description: `namespace of the vt entity`,
						Type:        smd.String,
					},
					{
						Name:        "entity",
						Optional:    true,
						Description: `VTEntity`,
						Type:        smd.Object,
						Properties: map[string]smd.Property{
							"name": {
								Description: ``,
								Type:        smd.String,
							},
							"terminalPath": {
								Description: ``,
								Type:        smd.String,
							},
							"attributes": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.VTAttributes",
								},
							},
							"template": {
								Description: ``,
								Type:        smd.Array,
								Items: map[string]string{
									"$ref": "#/definitions/mfd.TmplAttributes",
								},
							},
							"mode": {
								Description: ``,
								Type:        smd.String,
							},
						},
						Definitions: map[string]smd.Definition{
							"mfd.VTAttributes": {
								Type: "object",
								Properties: map[string]smd.Property{
									"name": {
										Description: ``,
										Type:        smd.String,
									},
									"attrName": {
										Description: ``,
										Type:        smd.String,
									},
									"searchName": {
										Description: ``,
										Type:        smd.String,
									},
									"summary": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"search": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"max": {
										Description: ``,
										Type:        smd.Integer,
									},
									"min": {
										Description: ``,
										Type:        smd.Integer,
									},
									"required": {
										Description: ``,
										Type:        smd.Boolean,
									},
									"validate": {
										Description: ``,
										Type:        smd.String,
									},
								},
							},
							"mfd.TmplAttributes": {
								Type: "object",
								Properties: map[string]smd.Property{
									"name": {
										Description: ``,
										Type:        smd.String,
									},
									"vtAttrName": {
										Description: ``,
										Type:        smd.String,
									},
									"list": {
										Description: `show in list`,
										Type:        smd.Boolean,
									},
									"fkOpts": {
										Description: `how to show fks`,
										Type:        smd.String,
									},
									"form": {
										Description: `show in object editor`,
										Type:        smd.String,
									},
									"search": {
										Description: `input type in search`,
										Type:        smd.String,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// Invoke is as generated code from zenrpc cmd
func (s XMLVTService) Invoke(ctx context.Context, method string, params json.RawMessage) zenrpc.Response {
	resp := zenrpc.Response{}
	var err error

	switch method {
	case RPC.XMLVTService.GenerateEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.GenerateEntity(args.Namespace, args.Entity))

	case RPC.XMLVTService.LoadEntity:
		var args = struct {
			Namespace string `json:"namespace"`
			Entity    string `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.LoadEntity(args.Namespace, args.Entity))

	case RPC.XMLVTService.SaveEntity:
		var args = struct {
			Namespace string        `json:"namespace"`
			Entity    *mfd.VTEntity `json:"entity"`
		}{}

		if zenrpc.IsArray(params) {
			if params, err = zenrpc.ConvertToObject([]string{"namespace", "entity"}, params); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		if len(params) > 0 {
			if err := json.Unmarshal(params, &args); err != nil {
				return zenrpc.NewResponseError(nil, zenrpc.InvalidParams, "", err.Error())
			}
		}

		resp.Set(s.SaveEntity(args.Namespace, args.Entity))

	default:
		resp = zenrpc.NewResponseError(nil, zenrpc.MethodNotFound, "", nil)
	}

	return resp
}
