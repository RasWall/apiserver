package controller

import (
	"github.com/gin-gonic/gin"
	// "github.com/shirou/gopsutil"
	"net/http"
)

// naming convention for Controllers
// [controllerClass]_[version]_[actionName]Controller(*gin.Context)

func GetCPUInfoController(c *gin.Context) {
	c.String(http.StatusOK, "cpuinfo")
}

func GetMemoryInfoController(c *gin.Context) {
	c.String(http.StatusOK, "meminfo")
}
