package config

import (
	"net/http"

	"github.com/cyjme/pkg"

	ConfigService "github.com/cyjme/service-center/config"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	req := request{}
	if err := pkg.ParseRequest(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ConfigService.Clear(req.Key)

	c.JSON(http.StatusOK, nil)
}
