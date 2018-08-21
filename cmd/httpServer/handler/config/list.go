package config

import (
	ConfigService "github.com/cyjme/service-center/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	configList := ConfigService.List("")

	c.JSON(http.StatusOK, configList)
}
