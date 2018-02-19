package main

import (
	"sandbox/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/*")

	r.POST("/payment", handler.Payment)
	r.POST("/payment/json", handler.PaymentJson)

	r.Run(":8080")
}
