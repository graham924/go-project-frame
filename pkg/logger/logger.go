package logger

import (
	"fmt"
	"go-project-frame/server/config"
	"go.uber.org/zap"
)

// Log 全局的Log对象
var Log *zap.Logger

// InitLogger 初始化Logger组件
func InitLogger() error {
	cfgLog := config.SysConfig.Log
	core, err := getZapCore(cfgLog)
	if err != nil {
		return err
	}
	// zap.AddCaller() 是一个可选配置，表示将调用函数信息记录到日志中
	Log = zap.New(core, zap.AddCaller())
	// 替换全局的 Logger 对象，全局 Logger 是一个单例对象，可以在程序中的任何地方使用 zap.L() 或 zap.S() 等全局函数来记录日志，这些函数会自动使用全局的 Logger 对象进行日志记录
	zap.ReplaceGlobals(Log)
	fmt.Println("init logger success")
	return nil
}
