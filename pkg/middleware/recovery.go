package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/logger"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

// Recovery 使用自定义日志库替换原有recover中间件，在发生 panic（宕机）时进行恢复和日志记录
// stack参数：是否需要记录 panic 的堆栈信息
func Recovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求结束之前，判断是否发生了panic
		defer func() {
			// 通过 recover() 函数获取到 panic 的错误信息 err
			if err := recover(); err != nil {
				// 判断是为网络连接发生的错误，这种不是需要保证紧急堆栈跟踪的真正情况
				var brokePipe bool
				// 断言err是否为*net.OpError类型
				if netErr, ok := err.(*net.OpError); ok {
					// 断言err是否为*os.SyscallError类型
					if sysErr, ok := netErr.Err.(*os.SyscallError); ok {
						// 如果属于调用失败，而且错误信息包含：broken pipe（连接管道中断）或 connection reset by peer（连接被重置）
						if strings.Contains(strings.ToLower(sysErr.Error()), "broken pipe") || strings.Contains(strings.ToLower(sysErr.Error()), "connection reset by peer") {
							brokePipe = true
						}
					}
				}

				// 将 HTTP 请求dump成一个字节数组，第二个参数表示是否包含request的budy
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				// 如果属于网络错误，则记录错误日志后，直接终止该请求，我们无法设置响应状态码
				if brokePipe {
					// 记录错误日志
					logger.Log.Error(ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 后面的这个注释，是告诉静态分析工具忽略这个检查
					ctx.Error(err.(error)) // nolint: errcheck

					// 终止当前请求的处理，并立即返回响应
					ctx.Abort()
					return
				}

				// 如果不是网络错误，则需要给前端设置响应状态码500
				if stack {
					logger.Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 终止当前请求的处理，并返回 HTTP 状态码 500
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		// 请求进入的时候，直接放行
		ctx.Next()
	}
}
