package mysql

import (
	"time"

	"github.com/qinsheng99/go-domain-web/domain/dp"
)

type Repo struct {
	Id           int64     `gorm:"column:id" json:"id"`
	RepoId       int64     `gorm:"column:repo_id" json:"repoId"`
	FullRepoName string    `gorm:"column:full_repo_name" json:"fullRepoName"`
	RepoName     string    `gorm:"column:repo_name" json:"repoName"`
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`
}

type RepoName struct {
	Id     int64  `gorm:"column:id" json:"id"`
	RepoId int64  `gorm:"column:repo_id" json:"repoId"`
	Name   string `gorm:"column:name" json:"name"`
}

func (r *RepoName) TableName() string {
	return "repo_name"
}

type RepoWith struct {
	Repo
	RepoNames []RepoName `json:"repo_names" gorm:"ForeignKey:repo_id;references:id"` //ForeignKey 连接表外键  references 对应主表的键
}

func (r *Repo) TableName() string {
	return "repo"
}

type RepoMapper interface {
	RepoNames(dp.Page, dp.Size, string) (data []Repo, err error)
	FindRepo(string) (data *Repo, err error)
	FindRepoWith(id int) (repo RepoWith, err error)
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

func (r *Repo) FindRepoName() string {
	Getmysqldb().Model(r).Where(r).First(r)

	return r.RepoName
}

func (r *Repo) Update() (err error) {
	err = Getmysqldb().Omit("create_time").Model(r).Updates(r).Error

	return
}

func (r *Repo) RepoNames(p dp.Page, s dp.Size, name string) (data []Repo, err error) {
	q := Getmysqldb().Model(r)
	if len(name) > 0 {
		q.Where("full_repo_name like ?", "%"+name+"%")
	}
	err = q.
		Order("full_repo_name asc").Limit(s.Size()).
		Offset((p.Page() - 1) * s.Size()).
		Find(&data).Error

	return
}

func (r *Repo) FindRepo(name string) (repo *Repo, err error) {
	repo = new(Repo)
	err = Getmysqldb().Model(r).
		Where("full_repo_name = ?", name).First(repo).Error

	return
}

func (r *Repo) FindRepoWith(id int) (repo RepoWith, err error) {
	err = Getmysqldb().Model(r).Where("id = ?", id).Preload("RepoNames").First(&repo).Error

	return
}
