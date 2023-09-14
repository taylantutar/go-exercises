package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, studentList)
}

func PostStudent(c *gin.Context) {
	var newStudent student
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	studentList = append(studentList, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func GetByIdStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, s := range studentList {
		if s.Id == id {
			c.JSON(http.StatusOK, s)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Student Not Found"})
	}
}

type student struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Age        int8   `json:"age"`
	Department string `json:"department"`
}

var studentList = []student{
	{Id: 1, Name: "Kaan", Surname: "Kaya", Age: 15, Department: "Sosyal"},
	{Id: 2, Name: "Melis", Surname: "Uçar", Age: 16, Department: "Fen"},
	{Id: 3, Name: "Murat", Surname: "Turan", Age: 15, Department: "Sosyal"},
	{Id: 4, Name: "Hakan", Surname: "Peker", Age: 15, Department: "Sosyal"},
}
