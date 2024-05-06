package main

import (
	"Twillo/Handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	// app.GET("/Echo", Handlers.EchoJSON)
	// app.GET("/RegisterTransaction", Handlers.RegisterTransaction)
	app.POST("/RegisterInstapay", Handlers.RegisterInstapayNew)
	app.Run(":4005")
}
