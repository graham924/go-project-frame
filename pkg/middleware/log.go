package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/logger"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录开始时间，请求path、query
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		// 调用 c.Next() 执行后续中间件和路由处理器的逻辑
		ctx.Next()

		cost := time.Since(start)
		logger.Log.Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("host", ctx.Request.Host),
			zap.Duration("cost", cost),
		)
	}
}
