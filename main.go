package main

import (
	"github.com/bots/api-imc/mails"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("archivo .env no se ha podido cargar de forma correcta")
	}

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/email/contact/send", mails.SendEmailTestResultController)
	}

	router.Run(os.Getenv("PORT_SERVER"))
}
