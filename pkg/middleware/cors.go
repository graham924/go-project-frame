package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 允许所有来源的请求访问该 API，即允许任意域名的请求。
		AllowAllOrigins: true,
		// 允许的HTTP方法列表
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		// 允许的请求头列表
		AllowHeaders: []string{"Content-Type", "Access-Token", "Authorization"},
		// 设置预请求的最大缓存时间
		MaxAge: 6 * time.Hour,
	})
}
