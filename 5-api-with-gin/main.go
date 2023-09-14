package main

import (
	"algorithm-exercises/5-api-with-gin/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting")
	router := gin.Default()

	router.GET("/student", handlers.GetStudents)
	router.POST("/student", handlers.PostStudent)

	router.Run("localhost:8081")
}
