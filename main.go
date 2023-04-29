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
	db := config.InitDb()

	userService := controllers.UserService{ DB: db }
	petService := controllers.PetService{ DB: db }
	speciesService := controllers.SpeciesService{ DB: db }
	breedService := controllers.BreedService{ DB: db }
	foodService := controllers.FoodService{ DB: db }
	medicineService := controllers.MedicineService{ DB: db }
	eventService := controllers.EventService{ DB: db }
	imageService := controllers.ImageService{ DB: db }

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", controllers.ServeSwaggerUI)
	router.GET("/ping", pingHandler)

	var api *gin.RouterGroup = router.Group("/api")
	{
		api.Use(middleware.TempUserAuth)
		var users *gin.RouterGroup = api.Group("/users")
		{
			users.GET("/", userService.GetUser)
			users.POST("/", userService.PostUser)
			users.DELETE("/", userService.DeleteUser)
		}
		var pets *gin.RouterGroup = api.Group("/pets")
		{
			pets.GET("/", petService.GetPets)
			pets.GET("/:petId", petService.GetPet)
			pets.POST("/", petService.PostPet)
			pets.PUT("/:petId", petService.PutPet)
			pets.DELETE("/:petId", petService.DeletePet)
		}
		var species *gin.RouterGroup = api.Group("/species")
		{
			species.GET("/:speciesId", speciesService.GetSpecies)
		}
		var breeds *gin.RouterGroup = api.Group("/breeds")
		{
			breeds.GET("/", breedService.GetBreeds)
			breeds.GET("/:breedId", breedService.GetBreed)
		}
		var foods *gin.RouterGroup = api.Group("/foods")
		{
			foods.GET("/:foodId", foodService.GetFood)
			foods.GET("/", foodService.GetFoods)
		}
		var medicines *gin.RouterGroup = api.Group("/medicines")
		{
			medicines.GET("/:medicineId", medicineService.GetMedicine)
			medicines.GET("/", medicineService.GetMedicines)
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
			events.GET("/", eventService.GetEvents)
			events.POST("/", eventService.PostEvent)
			events.PUT("/:eventId", eventService.PutEvent)
			events.DELETE("/:eventId", eventService.DeleteEvent)
		}
		var images *gin.RouterGroup = api.Group("/images")
		{
			images.GET("/pets/:petId", imageService.GetImagesByPet)
			images.GET("/users", imageService.GetImagesByUser)
			images.POST("/", imageService.PostImage)
			images.DELETE("/:imageId", imageService.DeleteImage)
		}
	}

	router.Run(":3000")
}
