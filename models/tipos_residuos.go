package models

import (
	"github.com/graphql-go/graphql"
)

type Tipos_Residuos struct {
	ID_Residuo  string `json:"id_residuo" gorm:"primary_key"`
	Descripcion string `json:"descripcion"`
}

var Tipos_ResiduosType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tipos_Residuos",
	Fields: graphql.Fields{
		"id_residuo": &graphql.Field{
			Type: graphql.String,
		},
		"descripcion": &graphql.Field{
			Type: graphql.String,
		},
	},
})
