package config

import (
	"net/http"

	ConfigService "github.com/cyjme/service-center/config"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, "key cannot be null")
		return
	}
	ConfigService.Clear(key)

	c.JSON(http.StatusOK, nil)
}
