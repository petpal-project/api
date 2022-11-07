package main

import (
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerMiddleware "github.com/go-openapi/runtime/middleware"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	// authClient := config.InitFirebase()
	config.InitDb()

	var r *gin.Engine = gin.Default()
	r.GET("/ping", pingHandler)

	opts := swaggerMiddleware.SwaggerUIOpts{SpecURL: "docs/openapi.yml"}
	sh := swaggerMiddleware.SwaggerUI(opts, nil)
	r.GET("/docs/*any", gin.WrapH(sh))

	var api *gin.RouterGroup = r.Group("/api")
	{
		// api.Use(middleware.TokenAuth(authClient, DB))
		api.Use(middleware.TempUserAuth)
		var users *gin.RouterGroup = api.Group("/users")
		{
			users.GET("/", controllers.GetUser)
			users.POST("/", controllers.PostUser)
			users.DELETE("/", controllers.DeleteUser)
		}
		var pets *gin.RouterGroup = api.Group("/pets")
		{
			pets.GET("/", controllers.GetPet)
			pets.POST("/", controllers.PostPet)
		}
		var species *gin.RouterGroup = api.Group("/species")
		{
			species.GET("/:speciesId", controllers.GetSpecies)
			species.POST("/")
		}
		var breeds *gin.RouterGroup = api.Group("/breeds")
		{
			breeds.GET("/")
			breeds.POST("/")
		}
		var foods *gin.RouterGroup = api.Group("/foods")
		{
			foods.GET("/")
			foods.POST("/")
		}
		var medicines *gin.RouterGroup = api.Group("/medicines")
		{
			medicines.GET("/")
			medicines.POST("/")
		}
	}

	r.Run(":3000")
}
