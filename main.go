package main

import (
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDb()

	pingService := controllers.PingService{}
	userService := controllers.UserService{DB: db}
	petService := controllers.PetService{DB: db}
	speciesService := controllers.SpeciesService{DB: db}
	breedService := controllers.BreedService{DB: db}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", controllers.ServeSwaggerUI)
	router.GET("/ping", pingService.Ping)

	api := router.Group("/api")
	{
		api.Use(middleware.TempUserAuth)
		users := api.Group("/users")
		{
			users.GET("/", userService.GetUser)
			users.POST("/", userService.PostUser)
			users.DELETE("/", userService.DeleteUser)
		}
		pets := api.Group("/pets")
		{
			pets.GET("/", petService.GetPets)
			pets.GET("/:petId", petService.GetPet)
			pets.POST("/", petService.PostPet)
			pets.PUT("/:petId", petService.PutPet)
			pets.DELETE("/:petId", petService.DeletePet)
		}
		species := api.Group("/species")
		{
			species.GET("/:speciesId", speciesService.GetSpecies)
		}
		breeds := api.Group("/breeds")
		{
			breeds.GET("/", breedService.GetBreeds)
			breeds.GET("/:breedId", breedService.GetBreed)
		}
	}

	router.Run(":3000")
}
