package handlers

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cf-neto/nettrace/internal/checker"
	"github.com/cf-neto/nettrace/internal/info"
	"github.com/cf-neto/nettrace/internal/metrics"
	"github.com/cf-neto/nettrace/internal/models"
)

func HealthHandler(ctx *gin.Context) {
	uptime := time.Since(info.StartTime).Round(time.Second)

	ctx.IndentedJSON(http.StatusOK, models.HealthResponse{
		Status: "ok",
		Uptime: uptime.String(),
	})
}

func VersionHandler(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, models.VersionResponse{
		Name:        info.Name,
		Version:     info.Version,
		Description: info.Description,
		Author:      info.Author,
		Repository:  info.Repository,
		Environment: info.Environment,
	})
}

func StatsHandler(ctx *gin.Context) {
	uptime := time.Since(info.StartTime).Round(time.Second)

	ctx.IndentedJSON(http.StatusOK, models.StatsResponse{
		TotalRequests: metrics.Total(),
		Uptime:        uptime.String(),
		StartedAt:     info.StartTime.Format(time.RFC3339),
	})
}

func CheckHandler(ctx *gin.Context) {
	target := ctx.Query("url")
	if target == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query param 'url' is required"})
		return
	}

	u, err := url.ParseRequestURI(target)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid url, use http:// or https://"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, checker.Check(target))
}