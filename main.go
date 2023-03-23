package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/latihanEksplorasi/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/movies", controllers.GetMovies)
	router.POST("/movies", controllers.InsertMovies)
	router.PUT("/movies/:id", controllers.UpdateMovies)
	router.DELETE("/movies/:id", controllers.DeleteMovies)
	router.Run("localhost:8080")
	fmt.Println("Serving localhost:8080")
	test := 7070
}
