package server

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"go-project-frame/pkg/logger"
	"go-project-frame/pkg/utils"
	"go-project-frame/server/config"
	"go-project-frame/server/initialize"
	"go-project-frame/server/options"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func NewServerCommand() *cobra.Command {
	opts, err := options.NewOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to initialize command options: %v", zap.Error(err))
	}
	cmd := &cobra.Command{
		Use:  "kubemanage-server",
		Long: "The kubemanage server controller is a daemon that embeds the core control loops.",
		Run: func(cmd *cobra.Command, args []string) {
			initServer(opts)
			runServer(opts)
		},
	}
	opts.BindConfigFileFlags(cmd)
	return cmd
}

func initServer(opts *options.Options) {
	if err := initialize.InitServer(opts); err != nil {
		fmt.Fprintf(os.Stderr, "unable to complete flags parse: %v", zap.Error(err))
		os.Exit(1)
	}
}

func runServer(opts *options.Options) {
	// 取出配置中的监听端口
	addr := config.SysConfig.Default.ListenAddr

	// 创建一个 http 的 web服务
	srv := &http.Server{
		Addr:    addr,
		Handler: opts.GinEngine, // 这表示：我们将使用 GinEngine 来处理 HTTP 请求
	}

	// Initializing the srv in a goroutine so that it won't block the graceful shutdown handling below
	// 开启一个goroutine，去执行server的端口监听操作，保证不会阻塞程序的正常终止
	go func() {
		// 记录 服务器已成功启动，并指定了监听地址
		logger.Log.Info("success", zap.String("starting kubemanage srv running on ", addr))
		// ListenAndServe 启动 HTTP 服务器，监听TCP连接，如果过程中有错误，并且不是 用户发起shutdown或Close所导致的ErrServerClosed，则终止服务器
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("failed to listen kubemanage server: ", zap.Error(err))
		}
	}()

	// 设置信号处理程序，返回一个通道quit。通过 <-quit 等待中断信号
	quit := utils.SetupSignalHandler()
	<-quit

	// 输出日志：服务器正在关闭
	logger.Log.Info("shutting kubemanage server down ...")

	// 创建一个带有超时时间的上下文，表示 服务器还有5s的时间处理业务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 程序退出时，调用cancel取消上下文
	defer cancel()

	// 优雅的关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		// 如果关闭过程出现错误，说明服务器已经被强制关闭了，打印错误信息，然后退出程序
		logger.Log.Fatal("kubemanage server forced to shutdown: ", zap.Error(err))
		os.Exit(1)
	}
	// 打印信息：服务器已关闭
	logger.Log.Info("kubemanage server exit successful")

}
