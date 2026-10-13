package metrics

import (
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var totalRequests atomic.Uint64 // the counter

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		totalRequests.Add(1)
		ctx.Next()
	}
}

func Total() uint64 {
	return totalRequests.Load()
}
