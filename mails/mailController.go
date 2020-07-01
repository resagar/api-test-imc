package mails

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type emailTestResult struct {
	Email string `json:"email" binding:"required"`
}

func SendEmailTestResultController(context *gin.Context) {
	var jsonData emailTestResult
	err := context.ShouldBind(&jsonData)
	if err != nil {
		log.Panic(err)
	}
	SendEmailTestResult(jsonData.Email, jsonData.Email, "Resultado de su Test")
	context.JSON(http.StatusOK, gin.H{"menssage": "Correo enviado con exito"})
}
