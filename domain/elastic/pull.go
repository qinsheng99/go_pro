package elastic

import (
	"context"

	"github.com/qinsheng99/go-domain-web/common/api"
	"github.com/qinsheng99/go-domain-web/domain"
)

type RepoPullImpl interface {
	Refresh(context.Context) (err error)
	PullList(api.RequestPull, context.Context) ([]domain.PullInfo, int64, error)
}
