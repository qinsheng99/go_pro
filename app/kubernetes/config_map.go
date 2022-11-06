package app

import (
	"context"
	"github.com/qinsheng99/go-domain-web/domain/kubernetes"
)

type ConfigMapServiceImpl interface {
	Create(ctx context.Context) (err error)
}

type configMapService struct {
	c kubernetes.ConfigMap
}

func NewConfigMapService(c kubernetes.ConfigMap) ConfigMapServiceImpl {
	return &configMapService{c: c}
}

func (c *configMapService) Create(ctx context.Context) (err error) {
	_, err = c.c.Create(ctx)
	return
}
