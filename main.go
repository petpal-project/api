package main

import (
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	config.InitDb()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", controllers.ServeSwaggerUI)
	router.GET("/ping", pingHandler)

	var api *gin.RouterGroup = router.Group("/api")
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
			pets.GET("/", controllers.GetPets)
			pets.GET("/:petId", controllers.GetPet)
			pets.POST("/", controllers.PostPet)
			pets.PUT("/:petId", controllers.PutPet)
			pets.DELETE("/:petId", controllers.DeletePet)
		}
		var species *gin.RouterGroup = api.Group("/species")
		{
			species.GET("/:speciesId", controllers.GetSpecies)
		}
		var breeds *gin.RouterGroup = api.Group("/breeds")
		{
			breeds.GET("/", controllers.GetBreeds)
			breeds.GET("/:breedId", controllers.GetBreed)
		}
		var foods *gin.RouterGroup = api.Group("/foods")
		{
			foods.GET("/:foodId", controllers.GetFood)
			foods.GET("/", controllers.GetFoods)
		}
		var medicines *gin.RouterGroup = api.Group("/medicines")
		{
			medicines.GET("/:medicineId", controllers.GetMedicine)
			medicines.GET("/", controllers.GetMedicines)
		}
		// for these two groups, could we want to have a new router group like
		// petMeals = *gin.RouterGroup = api.Group("/:mealId") ?
		var meals *gin.RouterGroup = api.Group("/meals")
		{
			meals.GET("/:mealId")
			meals.GET("/pet/:petId")
			meals.PUT("/:mealId")
			meals.POST("/")
			meals.DELETE("/:mealId")

		}
		var medications *gin.RouterGroup = api.Group("/medications")
		{
			medications.GET("/:medicationId")
			medications.GET("/pets/:petId")
			medications.PUT("/:medicationId")
			medications.POST("/")
			medications.DELETE("/:medicationId")
		}
		var events *gin.RouterGroup = api.Group("/events")
		{
			events.GET("/", controllers.GetEvents)
			events.POST("/", controllers.PostEvent)
			events.PUT("/:eventId", controllers.PutEvent)
			events.DELETE("/:eventId", controllers.DeleteEvent)
		}
		var images *gin.RouterGroup = api.Group("/images")
		{
			images.GET("/pets/:petId", controllers.GetImagesByPet)
			images.GET("/users", controllers.GetImagesByUser)
			images.POST("/", controllers.PostImage)
			images.DELETE("/:imageId", controllers.DeleteImage)
		}
	}

	router.Run(":3000")
}
