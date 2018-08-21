package router

import (
	"github.com/gin-gonic/gin"
	"github.com/cyjme/service-center/cmd/httpServer/handler/config"
)

func RegisterV1Router(r *gin.Engine) {
	serviceRouter := r.Group("/service")
	{
		serviceRouter.GET("")
	}

	configRouter := r.Group("/config")
	{
		configRouter.GET("", config.List)
		configRouter.DELETE("", config.Delete)
		configRouter.PUT("", config.Update)
	}
}
