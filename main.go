package main

import (
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDb()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", controllers.ServeSwaggerUI)
	router.GET("/ping", controllers.Ping)

	api := router.Group("/api")
	{
		api.Use(middleware.TempUserAuth)
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetUser(db))
			users.POST("/", controllers.PostUser(db))
			users.DELETE("/", controllers.DeleteUser(db))
		}
		pets := api.Group("/pets")
		{
			pets.GET("/", controllers.GetPets(db))
			pets.GET("/:petId", controllers.GetPet(db))
			pets.POST("/", controllers.PostPet(db))
			pets.PUT("/:petId", controllers.PutPet(db))
			pets.DELETE("/:petId", controllers.DeletePet(db))
		}
		species := api.Group("/species")
		{
			species.GET("/:speciesId", controllers.GetSpecies(db))
		}
		breeds := api.Group("/breeds")
		{
			breeds.GET("/", controllers.GetBreeds(db))
			breeds.GET("/:breedId", controllers.GetBreed(db))
		}
	}

	router.Run(":3000")
}
