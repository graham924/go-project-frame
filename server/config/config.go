package config

import (
	"go-project-frame/server/options"
	"os"
)

const (
	envConfigFileKey string = "KUBE_MANAGE_CONFIG"
)

var SysConfig *Config

type Config struct {
	Default DefaultOptions `mapstructure:"default"`
	Mysql   MysqlOptions   `mapstructure:"mysql"`
	Log     LogOptions     `mapstructure:"log"`
}

type DefaultOptions struct {
	ListenAddr     string `mapstructure:"listenAddr"`
	JWTSecret      string `mapstructure:"jwtSecret"`
	KubeConfigPath string `mapstructure:"kubeConfigPath"`
}

type MysqlOptions struct {
	// Host 主机名
	Host string `mapstructure:"host"`
	// Port 端口号
	Port string `mapstructure:"port"`
	// User 数据库用户名
	User string `mapstructure:"user"`
	// Password 数据库密码
	Password string `mapstructure:"password"`
	// Name 要操作的db数据库名称
	Name string `mapstructure:"name"`
	// MaxOpenConns 指定连接池中最大的打开连接数
	MaxOpenConns int `mapstructure:"maxOpenConns"`
	// MaxLifetime 指定连接的最大生命周期（秒）
	MaxLifetime int `mapstructure:"maxLifetime"`
	// MaxIdleConns 指定连接池中最大的空闲连接数
	MaxIdleConns int `mapstructure:"maxIdleConns"`
}

type LogOptions struct {
	// Level 日志级别
	Level string `mapstructure:"level"`
	// Filename 日志文件位置
	Filename string `mapstructure:"fileName"`
	// MaxSize 日志文件最大大小(MB)
	MaxSize int `mapstructure:"maxSize"`
	// MaxAge 保留旧日志文件的最大天数
	MaxAge int `mapstructure:"maxAge"`
	// MaxBackups 最大保留日志个数
	MaxBackups int `mapstructure:"maxBackups"`
}

func SysConfigParse(opts *options.Options) error {
	if len(opts.ConfigFile) == 0 {
		if envCfgFile := os.Getenv(envConfigFileKey); envCfgFile != "" {
			opts.ConfigFile = envCfgFile
		} else {
			opts.ConfigFile = options.DefaultConfigFile
		}
	}
	if err := BindingSysConfig(opts.ConfigFile); err != nil {
		return err
	}
	return nil
}
