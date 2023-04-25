package elasticsearch

import (
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/utils/gitee"
)

type pullRequestDO struct {
	Uuid        uuid.UUID      `json:"-"`
	Id          int            `json:"id"`
	Org         string         `json:"org"           description:"组织"`
	Ref         string         `json:"ref"           description:"目标分支"`
	Sig         string         `json:"sig"           description:"所属sig组"`
	Link        string         `json:"link"          description:"链接"`
	Repo        string         `json:"repo"          description:"仓库"`
	State       string         `json:"state"         description:"状态"`
	Title       string         `json:"title"         description:"标题"`
	Author      string         `json:"author"        description:"提交者"`
	CreatedAt   string         `json:"created_at"    description:"PR创建时间"`
	UpdatedAt   string         `json:"updated_at"    description:"PR更新时间"`
	Description string         `json:"description"   description:"描述"`
	Labels      pq.StringArray `json:"labels"        description:"标签"`
	Assignees   pq.StringArray `json:"assignees"     description:"指派者"`
	Draft       bool           `json:"draft"         description:"是否是草稿"`
	Mergeable   bool           `json:"mergeable"     description:"是否可合入"`
}

func toPullRequestDO(v *pullRequestDO, datum gitee.PullRequest, org, repo, htmlUrl string) {
	*v = pullRequestDO{
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
		Draft:       datum.GetDraft(),
		Mergeable:   datum.GetMergeable(),
	}
}
