package config

import (
	"net/http"

	ConfigService "github.com/cyjme/service-center/config"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	configList := ConfigService.List("")

	c.JSON(http.StatusOK, configList)
}
