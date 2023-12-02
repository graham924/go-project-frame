package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/consts"
	"go-project-frame/pkg/globalError"
	"go-project-frame/pkg/utils"
)

func Casbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果请求属于 非拦截路径，则直接放行
		if AlwaysAllowPath.Has(ctx.Request.URL.Path) {
			return
		}
		// 从上下文中获取前面 jwt 解析出的 claims
		_, exists := ctx.Get(consts.ClaimsContextKey)
		if !exists {
			utils.ResponseError(ctx, globalError.GetGlobalError(globalError.InternalServerError))
			ctx.Abort()
			return
		}

		// TODO
		ctx.Next()
	}
}
