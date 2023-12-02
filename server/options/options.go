package options

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-project-frame/dao"
	"gorm.io/gorm"
)

const (
	DefaultConfigFile string = "./config.yaml"
)

type Options struct {
	// GinEngine gin引擎对象
	GinEngine *gin.Engine
	// DB 数据库连接
	DB *gorm.DB
	// 数据库抽象工厂，包含所有数据库的操作接口
	Factory dao.ShareDaoFactory
	// 配置文件路径
	ConfigFile string
}

// NewOptions New one options with default ConfigFile
func NewOptions() (*Options, error) {
	return &Options{
		ConfigFile: DefaultConfigFile,
	}, nil
}

// BindConfigFileFlags bind config file flags
func (o *Options) BindConfigFileFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.ConfigFile, "configFilePath", "", "The location of the kubemanage configuration file")
}
