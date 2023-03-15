package main

import (
	"github.com/gin-gonic/gin"
	"github.com/latihanEksplorasi/controllers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.GET("/movies", controllers.GetMovies)
	router.POST("/movies",controllers.InsertMovies)
	router.PUT("/movies/:id", controllers.UpdateMovies)
	router.DELETE("/movies/:id", controllers.DeleteMovies)
	router.Run("localhost:8080")
}
