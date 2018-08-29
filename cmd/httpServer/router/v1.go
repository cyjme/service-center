package router

import (
	"github.com/cyjme/service-center/cmd/httpServer/handler/config"
	"github.com/cyjme/service-center/cmd/httpServer/handler/service"
	"github.com/gin-gonic/gin"
)

func RegisterV1Router(r *gin.Engine) {
	serviceRouter := r.Group("/service")
	{
		serviceRouter.GET("", service.List)
	}

	configRouter := r.Group("/config")
	{
		configRouter.GET("", config.List)
		configRouter.DELETE("/:key", config.Delete)
		configRouter.PUT("", config.Update)
	}
}
