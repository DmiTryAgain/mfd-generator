package vt

import (
	"html/template"

	"github.com/dizzyfool/genna/util"
	"github.com/vmkteam/mfd-generator/generators/model"
	base "github.com/vmkteam/mfd-generator/generators/repo"
	"github.com/vmkteam/mfd-generator/mfd"
)

// ServiceNamespaceData stores namespace info
type ServiceNamespaceData struct {
	Package      string
	ModelPackage string

	Name    string
	VarName string

	Entities []ServiceEntityData
}

// PackServiceNamespace packs mfd vt namespace to template data
func PackServiceNamespace(namespace *mfd.VTNamespace, options Options) ServiceNamespaceData {
	entities := make([]ServiceEntityData, len(namespace.Entities))
	for i, entity := range namespace.Entities {
		entities[i] = PackServiceEntity(*entity)
	}

	name := util.CamelCased(util.Sanitize(namespace.Name))

	return ServiceNamespaceData{
		Package:      options.Package,
		ModelPackage: options.ModelPackage,

		Name:    name,
		VarName: mfd.VarName(name),

		Entities: entities,
	}
}

func (tp ServiceNamespaceData) Raw(s string) template.HTML {
	return template.HTML(s)
}

// ServiceEntityData stores entity info
type ServiceEntityData struct {
	Name          string
	NamePlural    string
	VarName       string
	VarNamePlural string
	ShortVarName  string

	HasSortColumns bool
	SortColumns    []string

	PKs []base.PKPair

	HasAlias   bool
	PKSearches []base.PKPair
	AliasField string
	AliasArg   string

	HasRelations bool
	Relations    []ServiceRelationData
}

// PackServiceEntity packs mfd vt entity to template data
func PackServiceEntity(vtEntity mfd.VTEntity) ServiceEntityData {
	baseEntity := base.PackEntity(*vtEntity.Entity)

	var relations []ServiceRelationData
	var sortColumns []string
	for _, vtAttr := range vtEntity.Attributes {
		attr := vtAttr.Attribute
		if attr != nil && !attr.IsArray && vtAttr.Summary {
			sortColumns = append(sortColumns, attr.Name)
		}

		if attr != nil && attr.ForeignKey != "" && attr.ForeignEntity != nil {
			relations = append(relations, PackServiceRelationData(*vtAttr, *attr.ForeignEntity))
		}
	}

	// setting search for alias unique
	var pkSearches []base.PKPair
	var aliasField, aliasArg string
	for _, vtAttr := range vtEntity.Attributes {
		attr := vtAttr.Attribute

		if attr.DBName == "alias" {
			aliasField, aliasArg = attr.Name, attr.Name
		}

		if attr.PrimaryKey {
			search := vtEntity.Entity.SearchByAttrName(attr.Name, mfd.SearchNotEquals)
			if search == nil {
				continue
			}

			column := model.CustomSearchAttribute(*vtEntity.Entity, *search, model.Options{})
			pkSearches = append(pkSearches, base.PKPair{
				Field: attr.Name,
				Arg:   column.Name,
			})
		}
	}

	return ServiceEntityData{
		Name:          baseEntity.Name,
		NamePlural:    baseEntity.NamePlural,
		VarName:       baseEntity.VarName,
		VarNamePlural: baseEntity.VarNamePlural,
		ShortVarName:  mfd.ShortVarName(baseEntity.Name),

		HasSortColumns: len(sortColumns) > 0,
		SortColumns:    sortColumns,

		PKs: baseEntity.PKs,

		HasAlias:   aliasField != "",
		PKSearches: pkSearches,
		AliasField: aliasField,
		AliasArg:   aliasArg,

		HasRelations: len(relations) > 0,
		Relations:    relations,
	}
}

// ServiceRelationData stores relation info
type ServiceRelationData struct {
	Name      string
	NameSpace string
	FK        string
	PluralFK  string
	JSONName  string
	Nullable  bool
	IsArray   bool
}

// PackServiceRelationData packs mfd vt attribute to relation template data
func PackServiceRelationData(vtAttr mfd.VTAttribute, foreign mfd.Entity) ServiceRelationData {
	attr := vtAttr.Attribute

	baseRelation := model.PackRelation(*attr, model.Options{})

	return ServiceRelationData{
		Name:      attr.Name,
		JSONName:  mfd.JSONName(vtAttr.Name),
		FK:        baseRelation.ForeignEntity.Name,
		PluralFK:  mfd.MakePlural(util.CamelCased(baseRelation.ForeignEntity.Name)),
		NameSpace: foreign.Namespace,
		Nullable:  attr.Nullable(),
		IsArray:   attr.IsArray,
	}
}
