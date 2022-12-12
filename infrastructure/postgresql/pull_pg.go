package postgresql

import (
	"fmt"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/utils"
)

type Pull struct {
	Uuid        uuid.UUID      `gorm:"column:uuid;type:uuid" json:"-"`
	Id          int            `json:"id" gorm:"column:id;type:int8"`
	Org         string         `json:"org" description:"组织" gorm:"column:org;type:varchar(100)"`
	Repo        string         `json:"repo" description:"仓库" gorm:"column:repo;type:varchar(100)"`
	Ref         string         `json:"ref" description:"目标分支" gorm:"column:ref;type:varchar(100)"`
	Sig         string         `json:"sig" description:"所属sig组" gorm:"column:sig;type:varchar(100)"`
	Link        string         `json:"link" description:"链接" gorm:"column:link;type:varchar(255)"`
	State       string         `json:"state" description:"状态" gorm:"column:state;type:varchar(30)"`
	Author      string         `json:"author" description:"提交者" gorm:"column:author;type:varchar(50)"`
	Assignees   pq.StringArray `json:"assignees" description:"指派者" gorm:"column:assignees;type:text[];default:'{}'"`
	CreatedAt   string         `json:"created_at" description:"PR创建时间" gorm:"column:created_at;type:varchar(50)"`
	UpdatedAt   string         `json:"updated_at" description:"PR更新时间" gorm:"column:updated_at;type:varchar(50)"`
	Title       string         `json:"title" description:"标题" gorm:"column:title;type:text"`
	Description string         `json:"description" description:"描述" gorm:"column:description;type:text"`
	Labels      pq.StringArray `json:"labels" description:"标签" gorm:"column:labels;type:text[];default:'{}'"`
	Draft       bool           `json:"draft" description:"是否是草稿" gorm:"column:draft;type:bool"`
	Mergeable   bool           `json:"mergeable" description:"是否可合入" gorm:"column:mergeable;type:bool"`
}

const (
	Author    = "author"
	Assignees = "assignees"
	Labels    = "labels"
	Ref       = "ref"
	Sig       = "sig"
	Repo      = "repo"
)

func (Pull) TableName() string {
	return "pull"
}

type PullMapper interface {
	Exist(string) bool
	Insert(*Pull) error
	Update(*Pull) error
	FieldList(string, string, string, int, int) (int64, []string, error)
	FieldSliceList(string, string, int, int) (int64, []string, error)
	PullListForPG(api.RequestPull) ([]Pull, int64, error)
}

func NewPullMapper() PullMapper {
	return &Pull{}
}

func (p *Pull) Exist(link string) bool {
	var list []Pull
	_ = GetPostgresql().Model(p).Where("link = ?", link).Find(&list).Error
	if len(list) > 0 {
		return true
	}

	return false
}

func (p *Pull) Insert(pull *Pull) error {
	return GetPostgresql().Create(pull).Error
}

func (p *Pull) Update(pull *Pull) error {
	return GetPostgresql().Where("link = ?", pull.Link).Updates(pull).Error
}

func (p *Pull) FieldList(keyword, sig, field string, page, size int) (total int64, data []string, err error) {
	data = make([]string, 0)
	q := GetPostgresql().Model(p).Where("sig != ?", "Private").Where(field + " != ''")
	if len(keyword) > 0 {
		q.Where(field+" like ?", "%"+keyword+"%")
	}

	if len(sig) > 0 {
		q.Where("sig = ?", sig)
	}

	if err = q.Select(fmt.Sprintf("COUNT(DISTINCT %s)", field)).Count(&total).Error; err != nil || total == 0 {
		return
	}

	if field != Sig {
		q.Limit(size).Offset((page - 1) * size)
	}
	if err = q.Select("DISTINCT "+field).Order(field+" asc").Pluck(field, &data).Error; err != nil {
		return
	}

	return
}

func (p *Pull) FieldSliceList(keyword, field string, page, size int) (total int64, data []string, err error) {
	data = make([]string, 0)
	offset := int64((page - 1) * size)
	var list []Pull
	err = GetPostgresql().
		Where("sig != ?", "Private").
		Where(fmt.Sprintf("array_length(%s, 1) > 0", field)).
		Find(&list).Error
	if err != nil {
		return
	}
	var res []string
	switch field {
	case Labels:
		for _, pull := range list {
			res = append(res, pull.Labels...)
		}
	case Assignees:
		for _, pull := range list {
			res = append(res, pull.Assignees...)
		}
	}
	res = utils.FilterRepeat(res, keyword)
	total = int64(len(res))
	if len(res) == 0 || offset > total {
		return
	}
	sort.Strings(res)

	if total > offset && total < offset+int64(size) {
		data = res[offset:]
		return
	}
	data = res[offset : offset+int64(size)]
	return
}

func (p *Pull) PullListForPG(req api.RequestPull) (data []Pull, total int64, err error) {
	query := GetPostgresql().Model(p).Where("sig != 'Private'")
	if len(req.Label) > 0 {
		for _, s := range strings.Split(strings.ReplaceAll(req.Label, "，", ","), ",") {
			query.Where("? = ANY(labels)", s)
		}
	}
	if s := strings.Split(strings.ReplaceAll(req.State, "，", ","), ","); len(s) > 0 && len(req.State) > 0 {
		query.Where("state in ?", s)
	}

	if len(req.Org) > 0 {
		query.Where("org = ?", req.Org)
	}

	if len(req.Repo) > 0 {
		query.Where("repo = ?", req.Repo)
	}

	if len(req.Sig) > 0 {
		query.Where("sig = ?", req.Sig)
	}

	if len(req.Ref) > 0 {
		query.Where("ref = ?", req.Ref)
	}

	if len(req.Author) > 0 {
		query.Where("author = ?", req.Ref)
	}

	if len(req.Assignee) > 0 {
		query.Where("? = ANY(assignees)", req.Ref)
	}

	if len(req.Exclusion) > 0 {
		for _, s := range strings.Split(strings.ReplaceAll(req.Exclusion, "，", ","), ",") {
			query.Where("NOT ? = ANY(labels)", s)
		}
	}

	if len(req.Search) > 0 {
		query.Where("concat (repo, title, sig) like ?", "%"+req.Search+"%")
	}

	err = query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	query.Limit(req.PerPage).Offset((req.Page - 1) * req.PerPage).Order(fmt.Sprintf("%s %s", req.Sort, req.Direction))
	err = query.Debug().Find(&data).Error
	if err != nil {
		return nil, 0, err
	}
	return
}
