package elastic

import (
	"context"
	"strconv"

	"github.com/olivere/elastic/v7"
)

type Query struct {
	page, size      int
	includeSource   []string
	excludeSource   []string
	sort, sortField string
}

func NewQuery(page, size int, include, exclude []string, sort, sf string) *Query {
	return &Query{page: page, size: size, includeSource: include, excludeSource: exclude, sort: sort, sortField: sf}
}

type pullMapper struct {
	index string
}

func NewPullMapper(index string) pullMapper {
	return pullMapper{index: index}
}

func (p pullMapper) Insert(pulls map[int]interface{}, ctx context.Context) (err error) {
	for i, pl := range pulls {
		_, err = es.Index().Index(p.index).Id(strconv.Itoa(i)).BodyJson(pl).Do(ctx)
		if err != nil {
			return
		}
	}

	return
}

func (p pullMapper) UpdateColumn(doc interface{}, ctx context.Context, id string) (err error) {
	_, err = es.Update().Index(p.index).Id(id).Doc(doc).Do(ctx)

	return
}

func (p pullMapper) Exist(id string, ctx context.Context) (flag bool) {
	flag, _ = es.Exists().Index(p.index).Id(id).Do(ctx)

	return
}

func (p pullMapper) Update(pull interface{}, ctx context.Context, id string) (err error) {
	_, err = es.Index().Index(p.index).Id(id).BodyJson(pull).Do(ctx)

	return
}

func (p pullMapper) List(q elastic.Query, req *Query, ctx context.Context) (list []*elastic.SearchHit, total int64, err error) {
	search := p.baseSearch(q, req)
	r, err := search.Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	total = r.Hits.TotalHits.Value

	list = r.Hits.Hits

	return
}

func (p pullMapper) baseSearch(q elastic.Query, req *Query) *elastic.SearchService {
	search := es.Search().Index(p.index).Query(q).From((req.page - 1) * req.size).Size(req.size)
	if req.includeSource != nil || req.excludeSource != nil {
		search.FetchSourceContext(elastic.NewFetchSourceContext(true).Include(req.includeSource...).Exclude(req.excludeSource...))
	}

	if len(req.sort) > 0 && len(req.sortField) > 0 {
		search.Sort(req.sortField+".keyword", req.sort == "asc")
	}

	return search
}
