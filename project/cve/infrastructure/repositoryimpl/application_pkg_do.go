package repositoryimpl

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

type applicationPkgDO struct {
	Id          uuid.UUID `gorm:"column:uuid;type:uuid"    json:"-"`
	Org         string    `gorm:"column:org"               json:"-"`
	Repo        string    `gorm:"column:repo"              json:"-"`
	Assignee    string    `gorm:"column:assignee"          json:"assignee"`
	Version     string    `gorm:"column:version"           json:"-"`
	Platform    string    `gorm:"column:platform"          json:"-"`
	Community   string    `gorm:"column:community"         json:"-"`
	Milestone   string    `gorm:"column:milestone"         json:"milestone"`
	Decription  string    `gorm:"column:decription"        json:"decription"`
	PackageName string    `gorm:"column:package_name"      json:"-"`
	CreatedAt   string    `gorm:"column:created_at"        json:"-"`
	UpdatedAt   string    `gorm:"column:updated_at"        json:"updated_at"`
}

func (a applicationPkgImpl) toApplicationPkgDO(pkg *domain.ApplicationPackage) ([]applicationPkgDO, error) {
	var (
		res = make([]applicationPkgDO, 0)
		err error
	)
	for _, p := range pkg.Packages {
		do := applicationPkgDO{
			Id:          uuid.New(),
			Org:         pkg.Repository.Org,
			Repo:        pkg.Repository.Repo,
			Version:     p.Version,
			Platform:    pkg.Repository.Platform,
			Milestone:   p.Milestone,
			Community:   pkg.Repository.Community.Community(),
			Decription:  pkg.Repository.Desc.PackageDescription(),
			PackageName: p.Name.PackageName(),
			CreatedAt:   utils.Date(),
			UpdatedAt:   utils.Date(),
		}

		if p.Assignee != nil {
			do.Assignee = p.Assignee.Account()
		}

		if p.Id != "" {
			if do.Id, err = uuid.Parse(p.Id); err != nil {
				return nil, err
			}
		}

		res = append(res, do)
	}

	return res, err
}

func (a applicationPkgDO) toApplicationPkg() (v domain.ApplicationPackage, err error) {
	pkg := domain.Package{
		Id:        a.Id.String(),
		Version:   a.Version,
		Milestone: a.Milestone,
	}

	pkg.Assignee, _ = dp.NewAccount(a.Assignee)

	if pkg.Name, err = dp.NewPackageName(a.PackageName); err != nil {
		return
	}
	v.Packages = []domain.Package{pkg}
	v.Repository = domain.PackageRepository{
		Org:      a.Org,
		Repo:     a.Repo,
		Platform: a.Platform,
		Desc:     dp.NewDescription(a.Decription),
	}
	v.Repository.Community, err = dp.NewCommunity(a.Community)

	return
}

func (a applicationPkgDO) toMap() (map[string]interface{}, error) {
	var m map[string]interface{}

	v, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(v, &m)

	return m, err
}
