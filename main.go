package main

import (
	"todo/controllers"
	"todo/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un router con middleware default
	r := gin.Default()
	// Políticas de CORS (Intercambio de Recursos de Origen Cruzado)
	r.Use(cors.Default())
	// Apuntador a la base de datos
	db := models.SetupModels()
	defer db.Close() // Cerrar los recursos

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/graphql", controllers.Graphql)

	r.Run()
}
