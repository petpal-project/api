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
			events.GET("/:eventsId")
			events.GET("/pets/:petId")
			events.PUT("/:eventsId")
			events.POST("/")
			events.DELETE("/:eventsId")
		}
	}

	r.Run(":3000")
}
