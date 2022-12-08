package elasticsearch

import (
	"encoding/json"
	"strconv"

	"github.com/olivere/elastic/v7"
	"golang.org/x/net/context"
)

type Pull struct {
	Id          int    `json:"id"`
	Org         string `json:"org" description:"组织"`
	Repo        string `json:"repo" description:"仓库"`
	Ref         string `json:"ref" description:"目标分支"`
	Sig         string `json:"sig" description:"所属sig组"`
	Link        string `json:"link" description:"链接"`
	State       string `json:"state" description:"状态"`
	Author      string `json:"author" description:"提交者"`
	Assignees   string `json:"assignees" description:"指派者"`
	CreatedAt   string `json:"created_at" description:"PR创建时间"`
	UpdatedAt   string `json:"updated_at" description:"PR更新时间"`
	Title       string `json:"title" description:"标题"`
	Description string `json:"description" description:"描述"`
	Labels      string `json:"labels" description:"标签"`
}

type PullMapperImpl interface {
	Insert(pull []*Pull, ctx context.Context) (err error)
	PullList(elastic.Query, *Query, context.Context) ([]Pull, int64, error)
	UpdateColumn(interface{}, context.Context, string) error
	Update(*Pull, context.Context, string) error
	Exist(string, context.Context) bool
	AuthorList(elastic.Query, *Query, context.Context) ([]string, error)
}

type Query struct {
	page, size    int
	includeSource []string
	excludeSource []string
	sort          string
	sortField     string
}

func NewQuery(page, size int, include, exclude []string, sort string, sf string) *Query {
	return &Query{page: page, size: size, includeSource: include, excludeSource: exclude, sort: sort, sortField: sf}
}

type pullMapper struct {
	index string
}

func NewPullMapper(index string) PullMapperImpl {
	return pullMapper{index: index}
}

func (p pullMapper) Insert(pulls []*Pull, ctx context.Context) (err error) {
	for _, pl := range pulls {
		_, err = GetElasticsearch().Index().Index(p.index).Id(strconv.Itoa(pl.Id)).BodyJson(pl).Do(ctx)
		if err != nil {
			return
		}
	}
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

func (p pullMapper) AuthorList(q elastic.Query, req *Query, ctx context.Context) ([]string, error) {
	var list []string
	search := p.baseSearch(q, req)
	search.Collapse(elastic.NewCollapseBuilder("author.keyword"))
	r, err := search.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, hit := range r.Hits.Hits {
		a := struct {
			Author string `json:"author"`
		}{}
		err = json.Unmarshal(hit.Source, &a)
		if err != nil {
			return nil, err
		}
		list = append(list, a.Author)
	}
	return list, nil
}
