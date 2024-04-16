package main

import (
	"server/helper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register", helper.Register)

	r.POST("/login", helper.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
