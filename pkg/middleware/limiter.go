package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/globalError"
	"go-project-frame/pkg/utils"
	"golang.org/x/time/rate"
)

func Limiter() gin.HandlerFunc {
	/**
	 * 初始化一个限速器，每秒产生 1000 个令牌，桶的大小为 1000 个
	 * 初始化状态桶是满的
	 */
	limiter := rate.NewLimiter(1000, 1000)
	return func(ctx *gin.Context) {
		// 尝试去获取一个令牌
		if !limiter.Allow() {
			// 令牌获取失败，直接返回错误响应
			utils.ResponseError(ctx, globalError.NewGlobalError(
				globalError.InternalServerError,
				fmt.Errorf("系统繁忙，请稍后重试"),
			))
		}

	}
}
