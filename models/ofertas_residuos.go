package models

import (
	"github.com/graphql-go/graphql"
)

type Ofertas_Residuos struct {
	Id_Oferta  string `json:"id_oferta" gorm:"primary_key"`
	Id_Residuo string `json:"id_residuo"`
	Cantidad   int    `json:"cantidad"`
}

var Ofertas_ResiduosType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Residuos_Ofertas",
	Fields: graphql.Fields{
		"id_oferta": &graphql.Field{
			Type: graphql.String,
		},
		"id_residuo": &graphql.Field{
			Type: graphql.String,
		},
		"cantidad": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
