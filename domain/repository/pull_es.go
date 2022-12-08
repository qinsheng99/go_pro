package repository

import (
	"context"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
)

type RepoPullImpl interface {
	Refresh(context.Context) (err error)
	PullList(api.RequestPull, context.Context) ([]elasticsearch.Pull, int64, error)
	PullAuthor(api.RequestPull, context.Context) ([]string, error)
}
