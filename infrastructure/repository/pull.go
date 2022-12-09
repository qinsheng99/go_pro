package repository

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
	"github.com/olivere/elastic/v7"
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/domain/repository"
	"github.com/qinsheng99/go-domain-web/infrastructure/elasticsearch"
	"github.com/qinsheng99/go-domain-web/infrastructure/postgresql"
	"github.com/qinsheng99/go-domain-web/utils"
	_const "github.com/qinsheng99/go-domain-web/utils/const"
	"github.com/qinsheng99/go-domain-web/utils/gitee"
)

type repoPull struct {
	p   elasticsearch.PullMapperImpl
	pg  postgresql.PullMapper
	req utils.ReqImpl
}

func NewRepoPull(p elasticsearch.PullMapperImpl, req utils.ReqImpl, pg postgresql.PullMapper) repository.RepoPullImpl {
	return repoPull{p: p, req: req, pg: pg}
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

		pull := postgresql.Pull{
			Uuid:        uuid.New(),
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
		}

		if r.pg.Exist(pull.Link) {
			err = r.pg.Update(&pull)
			if err != nil {
				return err
			}
		} else {
			err = r.pg.Insert(&pull)
			if err != nil {
				return err
			}
		}

		err = r.p.InsertOne(&elasticsearch.Pull{Pull: pull}, ctx)
		if err != nil {
			return err
		}
	}
	return r.p.InsertMany(res, ctx)
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

func (r repoPull) PullFields(req api.RequestPull, field string) (int64, []string, error) {
	switch field {
	case _const.PullsAuthors:
		return r.pg.FieldList(req.Keyword, req.Sig, postgresql.Author, req.Page, req.PerPage)
	case _const.PullsAssignees:
		return r.pg.FieldSliceList(req.Keyword, postgresql.Assignees, req.Page, req.PerPage)
	case _const.PullsLabels:
		return r.pg.FieldSliceList(req.Keyword, postgresql.Labels, req.Page, req.PerPage)
	case _const.PullsRef:
		return r.pg.FieldList(req.Keyword, req.Sig, postgresql.Ref, req.Page, req.PerPage)
	case _const.PullsSig:
		return r.pg.FieldList(req.Keyword, req.Sig, postgresql.Sig, req.Page, req.PerPage)
	case _const.PullsRepos:
		return r.pg.FieldList(req.Keyword, req.Sig, postgresql.Repo, req.Page, req.PerPage)
	}

	return 0, nil, nil
}

func (r repoPull) PullAuthors(req api.RequestPull) (int64, []string, error) {
	return r.pg.FieldList(req.Keyword, req.Sig, postgresql.Author, req.Page, req.PerPage)
}

func (r repoPull) PullRef(req api.RequestPull) (int64, []string, error) {
	return r.pg.FieldList(req.Keyword, req.Sig, postgresql.Ref, req.Page, req.PerPage)
}
