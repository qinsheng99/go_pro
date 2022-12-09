package repository

import (
	"context"

	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
)

type RepoPullImpl interface {
	Refresh(context.Context) (err error)
	PullList(api.RequestPull, context.Context) ([]elasticsearch.Pull, int64, error)
	PullFields(api.RequestPull, string) (int64, []string, error)
	PullAuthors(api.RequestPull) (int64, []string, error)
	PullRef(api.RequestPull) (int64, []string, error)
}
