package service

import (
	"net/http"

	serviceService "github.com/cyjme/service-center/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	serviceList := serviceService.ListNodeByServiceName("")

	c.JSON(http.StatusOK, serviceList)
}
