package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project-frame/controller"
	v1 "go-project-frame/pkg/core/kubemanage/v1"
	"go-project-frame/pkg/database"
	"go-project-frame/pkg/jwt"
	"go-project-frame/pkg/kubernetes"
	"go-project-frame/pkg/logger"
	"go-project-frame/pkg/middleware"
	"go-project-frame/pkg/utils"
	"go-project-frame/server/config"
	"go-project-frame/server/options"
	"go.uber.org/zap"
	"os"
)

// InitServer init server
func InitServer(opts *options.Options) error {
	utils.PrintLogo()
	initGinEngine(opts)
	if err := configParse(opts); err != nil {
		fmt.Fprintf(os.Stderr, "unable to complete flags parse: %v\n", zap.Error(err))
		os.Exit(1)
	}
	if err := initLogger(); err != nil {
		return err
	}
	initJwt()
	if err := initDB(opts); err != nil {
		fmt.Fprintf(os.Stderr, "unable to complete db init: %v\n", zap.Error(err))
		os.Exit(1)
	}
	if err := initKubernetes(); err != nil {
		fmt.Fprintf(os.Stderr, "unable to complete kubernetes init: %v\n", zap.Error(err))
		os.Exit(1)
	}
	v1.SetupCoreService(opts)
	InstallRouters(opts)
	return nil
}

// init jwt
func initJwt() {
	jwt.JwtToken.InitJwt(config.SysConfig.Default.JWTSecret)
}

// initLogger init logger
func initLogger() error {
	return logger.InitLogger()
}

// initGinEngine init default engine
func initGinEngine(opts *options.Options) {
	opts.GinEngine = gin.Default()
}

// configParse completes all the required options
func configParse(opts *options.Options) error {
	return config.SysConfigParse(opts)
}

// initDB initialize db
func initDB(opts *options.Options) error {
	initDBService := database.InitDBService{}
	return initDBService.InitDB(opts)
}

// initKubernetes init kubernetes
func initKubernetes() error {
	return kubernetes.InitK8sClient()
}

// InstallRouters install routers
func InstallRouters(opts *options.Options) {
	apiGroup := opts.GinEngine.Group("api/")
	middleware.InstallMiddleware(apiGroup)
	controller.InstallRouter(apiGroup)
}
