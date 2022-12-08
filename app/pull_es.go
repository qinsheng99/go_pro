package app

import (
	"context"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
)

type pullService struct {
	pull repository.RepoPullImpl
}

type PullServiceImpl interface {
	Refresh(ctx context.Context) error
	PullList(req api.RequestPull, ctx context.Context) (list []elasticsearch.Pull, total int64, err error)
	PullAuthor(req api.RequestPull, ctx context.Context) ([]string, error)
}

func NewPullService(pull repository.RepoPullImpl) PullServiceImpl {
	return pullService{pull: pull}
}

func (p pullService) Refresh(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return p.pull.Refresh(ctx)
}

func (p pullService) PullList(req api.RequestPull, ctx context.Context) ([]elasticsearch.Pull, int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	return p.pull.PullList(req, ctx)
}

func (p pullService) PullAuthor(req api.RequestPull, ctx context.Context) ([]string, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	return p.pull.PullAuthor(req, ctx)
}
