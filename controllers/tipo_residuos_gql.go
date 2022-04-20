package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"todo/models"
)

func GetResiduos(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.Tipos_ResiduosType),
		Args: graphql.FieldConfigArgument{
			"id_residuo": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"descripcion": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var find models.Tipos_Residuos
			if id_residuo, isOK := params.Args["id_residuo"].(string); isOK == true {
				find.ID_Residuo = id_residuo
			}
			if descripcion, isOK := params.Args["descripcion"].(string); isOK == true {
				find.ID_Residuo = descripcion
			}

			var tipos []models.Tipos_Residuos
			db.Where(find).Find(&tipos)
			return tipos, nil
		},
	}
}
