package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/qinsheng99/go-domain-web/infrastructure/postgresql"
	"golang.org/x/net/context"
)

type Pull struct {
	postgresql.Pull
}

type PullMapperImpl interface {
	InsertMany(pull []*Pull, ctx context.Context) (err error)
	InsertOne(pull *Pull, ctx context.Context) (err error)
	PullList(elastic.Query, *Query, context.Context) ([]Pull, int64, error)
	UpdateColumn(interface{}, context.Context, string) error
	Update(*Pull, context.Context, string) error
	Exist(string, context.Context) bool
}

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

func NewPullMapper(index string) PullMapperImpl {
	return pullMapper{index: index}
}

func (p pullMapper) InsertMany(pulls []*Pull, ctx context.Context) (err error) {
	for _, pl := range pulls {
		_, err = GetElasticsearch().Index().Index(p.index).Id(strconv.Itoa(int(pl.Id))).BodyJson(pl).Do(ctx)
		if err != nil {
			return
		}
	}
	return
}

func (p pullMapper) InsertOne(pull *Pull, ctx context.Context) (err error) {
	_, err = GetElasticsearch().Index().Index(p.index).Id(strconv.Itoa(int(pull.Id))).BodyJson(pull).Do(ctx)
	return
}

func (p pullMapper) UpdateColumn(doc interface{}, ctx context.Context, id string) (err error) {
	_, err = GetElasticsearch().Update().Index(p.index).Id(id).Doc(doc).Do(ctx)
	return
}

func (p pullMapper) Exist(id string, ctx context.Context) (flag bool) {
	flag, _ = GetElasticsearch().Exists().Index(p.index).Id(id).Do(ctx)
	return
}

func (p pullMapper) Update(pull *Pull, ctx context.Context, id string) (err error) {
	_, err = GetElasticsearch().Index().Index(p.index).Id(id).BodyJson(pull).Do(ctx)
	return
}

func (p pullMapper) PullList(q elastic.Query, req *Query, ctx context.Context) (list []Pull, total int64, _ error) {
	search := p.baseSearch(q, req)
	r, err := search.Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	total = r.Hits.TotalHits.Value

	list = make([]Pull, 0)
	for _, hit := range r.Hits.Hits {
		var l Pull
		err = json.Unmarshal(hit.Source, &l)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, l)
	}

	return
}

func (p pullMapper) baseSearch(q elastic.Query, req *Query) *elastic.SearchService {
	search := GetElasticsearch().Search().Index(p.index).Query(q).From((req.page - 1) * req.size).Size(req.size)
	if req.includeSource != nil || req.excludeSource != nil {
		search.FetchSourceContext(elastic.NewFetchSourceContext(true).Include(req.includeSource...).Exclude(req.excludeSource...))
	}

	if len(req.sort) > 0 && len(req.sortField) > 0 {
		search.Sort(req.sortField+".keyword", req.sort == "asc")
	}
	return search
}
