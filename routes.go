package main

import (
	AuthV1Controller "github.com/RasWall/apiserver/controllers/v1/auth"
	SystemV1Controller "github.com/RasWall/apiserver/controllers/v1/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

var routes = Routes{
	Route{
		"Root",
		"GET",
		"/",
		false,
		func(c *gin.Context) {
			c.String(http.StatusOK, "root")
		},
	},
	Route{
		"System_v1 - CPU Info",
		"GET",
		"/v1/system/cpuinfo",
		true,
		SystemV1Controller.GetCPUInfoController,
	},
	Route{
		"System_v1 - CPU Info",
		"GET",
		"/v1/system/meminfo",
		true,
		SystemV1Controller.GetMemoryInfoController,
	},
	Route{
		"Auth_v1 - New Token",
		"GET",
		"/v1/auth/newtoken",
		false,
		AuthV1Controller.NewToken,
	},
}
