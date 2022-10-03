package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitializeRoutes(router *gin.Engine) {

	//Default routes
	InitializeMetricExporter()
	router.GET("/api/v1/health", Health)
	router.GET("/api/v1/metrics", gin.WrapH(promhttp.Handler()))
	router.NoRoute(NoRoute)
}
