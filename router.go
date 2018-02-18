package main

import (
	AuthV1Controller "github.com/RasWall/apiserver/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name       string
	Type       string
	Pattern    string
	Auth       bool
	HandleFunc func(*gin.Context)
}

type Routes []Route

func wrapAuth(handleFunc func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		AuthV1Controller.VerifyToken(c, handleFunc)
	}
}

func setupRoutes(r *gin.Engine) {

	for _, route := range routes {

		if route.Auth {
			route.HandleFunc = wrapAuth(route.HandleFunc)
		}

		switch route.Type {
		case "GET":
			r.GET(route.Pattern, route.HandleFunc)

		case "POST":
			r.POST(route.Pattern, route.HandleFunc)
		}
	}
}
