package main

import (
	"main/controllers"
	"main/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Static("/frontend", "./frontend")

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	r.GET("/horses", func(c *gin.Context) {
		type Horse struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
			Coat string `json:"coat"`
		}

		horses := []Horse{
			{"Dancer", 2, "chestnut"},
			{"Bingo", 8, "bay"},
			{"Pharlap", 4, "grey"},
		}

		c.JSON(200, horses)
	})

	r.POST("/horses", controllers.HorseCreate)

	r.Run() // listen and serve on PORT in environment variables.
}
