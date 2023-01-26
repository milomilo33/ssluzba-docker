package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"uns/database"
	"uns/models"

	"github.com/gin-gonic/gin"
)

var RequestCounter int

func RegisterStudent(context *gin.Context) {
	RequestCounter++
	fmt.Println("Current request counter: " + strconv.Itoa(RequestCounter))

	var student models.Student
	if err := context.ShouldBindJSON(&student); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&student)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.Status(http.StatusCreated)
}

func RegisterProfesor(context *gin.Context) {
	RequestCounter++
	fmt.Println("Current request counter: " + strconv.Itoa(RequestCounter))

	var profesor models.Profesor
	if err := context.ShouldBindJSON(&profesor); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&profesor)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.Status(http.StatusCreated)
}
