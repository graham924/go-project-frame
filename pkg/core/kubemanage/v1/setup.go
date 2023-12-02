package v1

import (
	"go-project-frame/pkg/core/kubemanage/v1/service"
	"go-project-frame/server/config"
	"go-project-frame/server/options"
)

var CoreV1 service.CoreService

func SetupCoreService(opts *options.Options) {
	CoreV1 = service.New(config.SysConfig, opts.Factory)
}
