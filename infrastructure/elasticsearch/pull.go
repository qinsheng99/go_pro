package elasticsearch

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/olivere/elastic/v7"

	"github.com/qinsheng99/go-domain-web/common/api"
	elasticlocal "github.com/qinsheng99/go-domain-web/common/infrastructure/elastic"
	"github.com/qinsheng99/go-domain-web/domain"
	elastic2 "github.com/qinsheng99/go-domain-web/domain/elastic"
	"github.com/qinsheng99/go-domain-web/utils"
	"github.com/qinsheng99/go-domain-web/utils/gitee"
)

type repoPull struct {
	cli esDao
	req utils.ReqImpl
}

func NewRepoPull(cli esDao, req utils.ReqImpl) elastic2.RepoPullImpl {
	return repoPull{cli: cli, req: req}
}

func (r repoPull) Refresh(ctx context.Context) error {
	url := "https://gitee.com/api/v5/enterprise/open_euler/pull_requests?state=all&sort=" +
		"created&direction=desc&page=1&per_page=100&access_token=xx"

	var data []gitee.PullRequest
	_, err := r.req.CustomRequest(url, "GET", nil, nil, nil, false, &data)
	if err != nil {
		return err
	}

	var res = make(map[int]interface{})
	for _, datum := range data {
		htmlUrl := datum.GetHtml()
		org := strings.Split(htmlUrl, "/")[3]
		if org != "src-openeuler" && org != "openeuler" {
			continue
		}
		repo := strings.Split(htmlUrl, "/")[4]

		var do pullRequestDO
		toPullRequestDO(&do, datum, org, repo, htmlUrl)

		res[datum.Id] = &do
	}
	return r.cli.Insert(res, context.Background())
}

func (r repoPull) PullList(req api.RequestPull, ctx context.Context) ([]domain.PullInfo, int64, error) {
	q := elastic.NewBoolQuery()

	if len(req.Label) > 0 {
		l := elastic.NewBoolQuery()
		for _, s := range strings.Split(strings.ReplaceAll(req.Label, "，", ","), ",") {
			l.Must(elastic.NewMatchPhraseQuery("labels", s))
		}
		q.Must(l)
	}
	//if l := utils.StrSliceToInterface(strings.Split(strings.ReplaceAll(req.Label, "，", ","), ",")); len(l) > 0 {
	//	q.Must(elastic.NewTermsQuery("labels", l...))
	//}

	if s := utils.StrSliceToInterface(strings.Split(strings.ReplaceAll(req.State, "，", ","), ",")); len(s) > 0 {
		q.Must(elastic.NewTermsQuery("state", s...))
	}

	if len(req.Org) > 0 {
		q.Must(elastic.NewTermQuery("org", req.Org))
	}

	if len(req.Repo) > 0 {
		q.Must(elastic.NewTermQuery("repo", req.Repo))
	}

	if len(req.Sig) > 0 {
		q.Must(elastic.NewTermQuery("sig", req.Sig))
	}

	if len(req.Ref) > 0 {
		q.Must(elastic.NewTermQuery("ref", req.Ref))
	}

	if len(req.Author) > 0 {
		q.Must(elastic.NewTermQuery("author", req.Author))
	}

	if len(req.Assignee) > 0 {
		q.Must(elastic.NewMatchPhraseQuery("assignees", req.Assignee))
	}

	if len(req.Exclusion) > 0 {
		e := elastic.NewBoolQuery()
		for _, s := range strings.Split(strings.ReplaceAll(req.Exclusion, "，", ","), ",") {
			e.Must(elastic.NewMatchPhraseQuery("labels", s))
		}
		q.MustNot(e)
	}

	if len(req.Search) > 0 {
		s := elastic.NewBoolQuery()
		for _, field := range []string{"repo", "title", "sig"} {
			s.Should(elastic.NewWildcardQuery(field, "*"+req.Search+"*"))
		}
		q.Must(s)
	}
	v, total, err := r.cli.List(
		q,
		elasticlocal.NewQuery(req.Page, req.PerPage, nil, nil, req.Direction, req.Sort), ctx,
	)
	if err != nil {
		return nil, 0, err
	}

	var res = make([]domain.PullInfo, total)
	for i, hit := range v {
		var p domain.PullInfo
		if err = json.Unmarshal(hit.Source, &p); err != nil {
			return nil, 0, err
		}

		res[i] = p
	}

	return res, total, nil
}
