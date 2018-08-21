package service

import (
	serviceService "github.com/cyjme/service-center/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	serviceList := serviceService.ListNodeByServiceName("")

	c.JSON(http.StatusOK, serviceList)
}
