package main

import (
	"main/controllers"
	"main/initializers"
	"main/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Horse{})
	//models.CreateSourceHorses()

	r := gin.Default()

	r.Static("/frontend", "./frontend")

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	r.GET("/horses", func(c *gin.Context) {
		name := c.Query("name")
		id := c.Query("id")

		var horse models.Horse

		if name != "" {
			result := initializers.DB.Where("name = ?", name).First(&horse)
			if result.Error != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch horse"})
				return
			}
		}

		if id != "" {
			result := initializers.DB.First(&horse, id)
			if result.Error != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch horse"})
				return
			}
		}

		c.JSON(200, &horse)

	})

	r.POST("/horses", controllers.HorseCreate)

	r.Run() // listen and serve on PORT in environment variables.
}
