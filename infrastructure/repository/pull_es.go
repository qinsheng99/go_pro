package repository

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
	"github.com/qinsheng99/go-domain-web/utils"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
	"github.com/qinsheng99/go-domain-web/utils/gitee"
)

const (
	author = "author"
)

type repoPull struct {
	p   elasticsearch.PullMapperImpl
	req utils.ReqImpl
}

func NewRepoPull(p elasticsearch.PullMapperImpl, req utils.ReqImpl) repository.RepoPullImpl {
	return repoPull{p: p, req: req}
}

func (r repoPull) Refresh(ctx context.Context) error {
	url := "https://gitee.com/api/v5/enterprise/open_euler/pull_requests?state=all&sort=" +
		"created&direction=desc&page=1&per_page=100&access_token=xx"

	var data []gitee.PullRequest
	_, err := r.req.CustomRequest(url, "GET", nil, nil, nil, false, &data)
	if err != nil {
		return err
	}

	var res []*elasticsearch.Pull
	for _, datum := range data {
		htmlUrl := datum.GetHtml()
		org := strings.Split(htmlUrl, "/")[3]
		if org != "src-openeuler" && org != "openeuler" {
			continue
		}
		repo := strings.Split(htmlUrl, "/")[4]

		res = append(res, &elasticsearch.Pull{
			Id:          datum.Id,
			Org:         org,
			Repo:        org + "/" + repo,
			Ref:         datum.GetRef(),
			Sig:         "",
			Link:        htmlUrl,
			State:       datum.GetState(),
			Author:      datum.GetLogin(),
			Assignees:   datum.GetAssignessName(),
			CreatedAt:   datum.GetCreate(),
			UpdatedAt:   datum.GetUpdate(),
			Title:       datum.GetTitle(),
			Description: base64.StdEncoding.EncodeToString([]byte(datum.GetBody())),
			Labels:      datum.GetLabelName(),
		})
	}
	return r.p.Insert(res, ctx)
}

func (r repoPull) PullList(req api.RequestPull, ctx context.Context) ([]elasticsearch.Pull, int64, error) {
	q := elastic.NewBoolQuery()

	if len(req.Label) > 0 {
		l := elastic.NewBoolQuery()
		for _, s := range strings.Split(strings.ReplaceAll(req.Label, "，", ","), ",") {
			l.Should(elastic.NewMatchPhraseQuery("labels", s))
		}
		q.Must(l)
	}

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
		q = q.Must(elastic.NewTermQuery("author", req.Author))
	}

	if len(req.Assignee) > 0 {
		q = q.Must(elastic.NewMatchPhraseQuery("assignees", req.Assignee))
	}

	if len(req.Exclusion) > 0 {
		e := elastic.NewBoolQuery()
		for _, s := range strings.Split(strings.ReplaceAll(req.Exclusion, "，", ","), ",") {
			e.Should(elastic.NewMatchPhraseQuery("labels", s))
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

	return r.p.PullList(
		q,
		elasticsearch.NewQuery(req.Page, req.PerPage, nil, nil, req.Direction, req.Sort),
		ctx,
	)
}

func (r repoPull) PullAuthor(req api.RequestPull, ctx context.Context) ([]string, error) {
	q := elastic.NewBoolQuery()
	if len(req.Keyword) > 0 {
		q.Must(elastic.NewWildcardQuery(author, "*"+req.Keyword+"*"))
	}

	return r.p.AuthorList(q, elasticsearch.NewQuery(req.Page, req.PerPage, []string{author}, nil, _const.Asc, author), ctx)
}
