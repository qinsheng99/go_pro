package app

import (
	"context"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/domain/elastic"
)

type pullService struct {
	pull elastic.RepoPullImpl
}

type PullServiceImpl interface {
	Refresh(ctx context.Context) error
	PullList(req api.RequestPull, ctx context.Context) (pullRequestDTO, error)
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

func (p pullService) PullList(req api.RequestPull, ctx context.Context) (pullRequestDTO, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	list, total, err := p.pull.PullList(req, ctx)
	if err != nil {
		return pullRequestDTO{}, err
	}

	return toPullRequestDTO(list, total), nil
}
