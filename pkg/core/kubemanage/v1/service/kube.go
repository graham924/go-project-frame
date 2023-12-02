package service

type KubeGetter interface {
	Kube() KubeInterface
}

type KubeInterface interface {
}

type kube struct {
	coreSvc *coreService
}

func NewKube(c *coreService) KubeInterface {
	return &kube{
		coreSvc: c,
	}
}
