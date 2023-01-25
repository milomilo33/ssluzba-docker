package controllers

import (
	"net/http"
	"uns/database"
	"uns/models"

	"github.com/gin-gonic/gin"
)

func RegisterStudent(context *gin.Context) {
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
