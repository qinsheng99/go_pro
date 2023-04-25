package app

import (
	"context"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/elastic"
)

type pullService struct {
	pull elastic.RepoPullImpl
}

type PullServiceImpl interface {
	Refresh(ctx context.Context) error
	PullList(req api.RequestPull, ctx context.Context) (list []domain.PullInfo, total int64, err error)
}

func NewPullService(pull elastic.RepoPullImpl) PullServiceImpl {
	return pullService{pull: pull}
}

func (p pullService) Refresh(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return p.pull.Refresh(ctx)
}

func (p pullService) PullList(req api.RequestPull, ctx context.Context) ([]pullRequestDTO, int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	return p.pull.PullList(req, ctx)
}
