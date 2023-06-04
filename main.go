package main

import (
	"log"
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db := config.InitDb()

	// Initialize Auth0 token validator
	validator, err := config.NewAuth0JWTValidator()
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", controllers.ServeSwaggerUI)
	router.GET("/ping", controllers.Ping)

	api := router.Group("/api")
	{
		api.Use(middleware.EnsureValidToken(validator))
		pets := api.Group("/pets")
		{
			pets.GET("/", controllers.GetPets(db))
			pets.POST("/", controllers.PostPet(db))
			pets.GET("/:petId", controllers.GetPet(db))
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
