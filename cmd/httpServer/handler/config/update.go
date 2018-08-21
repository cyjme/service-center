package config

import (
	ConfigService "github.com/cyjme/service-center/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {
	request := request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parse request error"})
	}

	ConfigService.Set(request.Key, request.Value)

	c.JSON(http.StatusOK, nil)
}
