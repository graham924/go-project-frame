package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/consts"
	"go-project-frame/pkg/globalError"
	"go-project-frame/pkg/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 如果请求属于 非拦截路径，则直接放行
		if AlwaysAllowPath.Has(ctx.Request.URL.Path) {
			// 在 Gin 中间件中直接使用 return 语句，相当于提前结束当前中间件的执行，并将控制权返回给请求处理链的下一个中间件或处理函数。
			// 这意味着后续的中间件或处理函数将不会被执行
			// TODO 直接 return 和 c.Next 有什么区别
			return
		}

		//// 如果是登录请求，直接放行，无需后续验证
		//if len(ctx.Request.URL.String()) == 15 && ctx.Request.URL.String() == "/api/user/login" {
		//	ctx.Next()
		//	return
		//}

		// 从上下文中获取token签名
		claims, err := utils.GetClaims(ctx)
		if err != nil {
			utils.ResponseError(ctx, globalError.NewGlobalError(globalError.AuthorizationError, err))
			ctx.Abort()
			return
		}

		// 签名存入上下文，供后续使用
		ctx.Set(consts.ClaimsContextKey, claims)
		ctx.Next()
	}
}
