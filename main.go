package main

import (
	"net/http"
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDb()

	pingService := &controllers.PingService{}
	userService := &controllers.UserService{DB: db}
	petService := &controllers.PetService{DB: db}
	speciesService := &controllers.SpeciesService{DB: db}
	breedService := &controllers.BreedService{DB: db}
	foodService := &controllers.FoodService{DB: db}
	medicineService := &controllers.MedicineService{DB: db}
	eventService := &controllers.EventService{DB: db}
	imageService := &controllers.ImageService{DB: db}

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
		foods := api.Group("/foods")
		{
			foods.GET("/:foodId", foodService.GetFood)
			foods.GET("/", foodService.GetFoods)
		}
		medicines := api.Group("/medicines")
		{
			medicines.GET("/:medicineId", medicineService.GetMedicine)
			medicines.GET("/", medicineService.GetMedicines)
		}
		// for these two groups, could we want to have a new router group like
		// petMeals = *gin.RouterGroup = api.Group("/:mealId") ?
		meals := api.Group("/meals")
		{
			meals.GET("/:mealId")
			meals.GET("/pet/:petId")
			meals.PUT("/:mealId")
			meals.POST("/")
			meals.DELETE("/:mealId")

		}
		medications := api.Group("/medications")
		{
			medications.GET("/:medicationId")
			medications.GET("/pets/:petId")
			medications.PUT("/:medicationId")
			medications.POST("/")
			medications.DELETE("/:medicationId")
		}
		events := api.Group("/events")
		{
			events.GET("/", eventService.GetEvents)
			events.POST("/", eventService.PostEvent)
			events.PUT("/:eventId", eventService.PutEvent)
			events.DELETE("/:eventId", eventService.DeleteEvent)
		}
		images := api.Group("/images")
		{
			images.GET("/pets/:petId", imageService.GetImagesByPet)
			images.GET("/users", imageService.GetImagesByUser)
			images.POST("/", imageService.PostImage)
			images.DELETE("/:imageId", imageService.DeleteImage)
		}
	}

	http.ListenAndServe(":3000", router)
}
