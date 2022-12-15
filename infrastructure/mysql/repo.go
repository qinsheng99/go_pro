package mysql

import (
	"github.com/qinsheng99/go-domain-web/api"
	"time"
)

type Repo struct {
	Id         int64     `gorm:"column:id" json:"id"`
	RepoId     int64     `gorm:"column:repo_id" json:"repoId"`
	Repo       string    `gorm:"column:repo" json:"repo"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (r *Repo) TableName() string {
	return "repo"
}

type RepoMapper interface {
	RepoNames(api.Pages) (data []string, err error)
	FindRepo(string) (data *Repo, err error)
}

func NewRepoMapper() RepoMapper {
	return &Repo{}
}

func (r *Repo) Insert() (err error) {
	err = Getmysqldb().Model(r).Create(r).Error
	return
}

func (r *Repo) Exist() bool {
	err := Getmysqldb().Model(r).Select("id").Where("repo_id = ?", r.RepoId).First(r).Error
	if err != nil {
		return false
	}

	return true
}

func (r *Repo) Update() (err error) {
	err = Getmysqldb().Omit("create_time").Model(r).Updates(r).Error
	return
}

func (r *Repo) RepoNames(p api.Pages) (data []string, err error) {
	p.SetDefault()
	err = Getmysqldb().Model(r).
		Order("repo asc").Limit(p.Size).
		Offset((p.Page-1)*p.Size).
		Pluck("repo", &data).Error
	return
}

func (r *Repo) FindRepo(name string) (repo *Repo, err error) {
	repo = new(Repo)
	err = Getmysqldb().Model(r).
		Where("repo = ?", name).First(repo).Error
	return
}
