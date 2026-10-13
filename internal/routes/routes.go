package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/cf-neto/nettrace/internal/handlers"
)

func Setup(r *gin.Engine) {
	r.GET("/health", handlers.HealthHandler)
	r.GET("/version", handlers.VersionHandler)
	r.GET("/stats", handlers.StatsHandler)
	r.GET("/check", handlers.CheckHandler)
}
