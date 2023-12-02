package service

type SystemGetter interface {
	System() SystemInterface
}

type SystemInterface interface {
}

type system struct {
	coreSvc *coreService
}

func NewSystem(c *coreService) SystemInterface {
	return &system{
		coreSvc: c,
	}
}
