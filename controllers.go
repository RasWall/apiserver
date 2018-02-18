package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// naming convention for Controllers
// [controllerClass]_[version]_[actionName]Controller(*gin.Context)

func System_v1_GetCPUInfoController(c *gin.Context) {
	c.String(http.StatusOK, "cpuinfo")
}

func System_v1_GetMemoryInfoController(c *gin.Context) {
	c.String(http.StatusOK, "meminfo")
}
