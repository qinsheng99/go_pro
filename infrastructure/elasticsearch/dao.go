package elasticsearch

import (
	"context"

	"github.com/olivere/elastic/v7"

	elasticlocal "github.com/qinsheng99/go-domain-web/common/infrastructure/elastic"
)

type esDao interface {
	Insert(pulls map[int]interface{}, ctx context.Context) (err error)
	UpdateColumn(doc interface{}, ctx context.Context, id string) (err error)
	Exist(id string, ctx context.Context) (flag bool)
	Update(pull interface{}, ctx context.Context, id string) (err error)
	List(q elastic.Query, req *elasticlocal.Query, ctx context.Context) (list []*elastic.SearchHit, total int64, _ error)
}
