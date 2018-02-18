package main

import (
	"github.com/gin-gonic/gin"
)

var err error

func main() {

	r := gin.Default()
	setupRoutes(r)

	r.Run(":8080")
}
