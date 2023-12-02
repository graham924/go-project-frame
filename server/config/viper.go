package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func BindingSysConfig(configFile string) error {
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config failed: %v", zap.Error(err))
	}
	if err := v.Unmarshal(&SysConfig); err != nil {
		return fmt.Errorf("config unmarshal failed: %v", zap.Error(err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed, sys config reload")
		if err := viper.Unmarshal(&SysConfig); err != nil {
			fmt.Errorf("sys config reload failed: %v", zap.Error(err))
		}
	})
	return nil
}
