package service

import (
	"go-project-frame/dao"
	"go-project-frame/server/config"
)

// CoreService service的核心入口
type CoreService interface {
	SystemGetter
	KubeGetter
}

func New(config *config.Config, f dao.ShareDaoFactory) CoreService {
	return &coreService{
		Cfg:     config,
		Factory: f,
	}
}

type coreService struct {
	Cfg     *config.Config
	Factory dao.ShareDaoFactory
}

func (c *coreService) System() SystemInterface {
	return NewSystem(c)
}

func (c *coreService) Kube() KubeInterface {
	return NewKube(c)
}
