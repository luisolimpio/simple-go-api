package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luisolimpio/api-go-gin/database"
	"github.com/luisolimpio/api-go-gin/models"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func PostStudent(c *gin.Context) {
	var newStudent models.Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

		return
	}

	database.DB.Create(&newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func GetStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func PatchStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.Delete(&student, id)
	c.JSON(http.StatusNoContent, gin.H{})
}

func GetStudentByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	var student models.Student

	database.DB.Where(&models.Student{Cpf: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}
