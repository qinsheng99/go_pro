package app

import (
	"context"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
	"github.com/qinsheng99/go-domain-web/infrastructure/postgresql"
)

type pullService struct {
	pull repository.RepoPullImpl
}

type PullServiceImpl interface {
	Refresh(ctx context.Context) error
	PullList(req api.RequestPull, ctx context.Context) (list []elasticsearch.Pull, total int64, err error)
	PullFields(api.RequestPull, string) (int64, []string, error)
	PullAuthors(api.RequestPull) (int64, []string, error)
	PullRef(api.RequestPull) (int64, []string, error)
	PullListPg(req api.RequestPull) ([]postgresql.Pull, int64, error)
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

func (p pullService) PullFields(req api.RequestPull, field string) (int64, []string, error) {
	return p.pull.PullFields(req, field)
}

func (p pullService) PullAuthors(req api.RequestPull) (int64, []string, error) {
	return p.pull.PullAuthors(req)
}

func (p pullService) PullRef(req api.RequestPull) (int64, []string, error) {
	return p.pull.PullRef(req)
}

func (p pullService) PullListPg(req api.RequestPull) ([]postgresql.Pull, int64, error) {
	return p.pull.PullListPg(req)
}
