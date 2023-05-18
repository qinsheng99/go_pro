package repositoryimpl

import (
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

type basePkgDO struct {
	Id          uuid.UUID      `gorm:"column:uuid;type:uuid"                    json:"-"`
	Org         string         `gorm:"column:org"                               json:"-"`
	Repo        string         `gorm:"column:repo"                              json:"repo"`
	Platform    string         `gorm:"column:platform"                          json:"-"`
	Community   string         `gorm:"column:community"                         json:"-"`
	Decription  string         `gorm:"column:decription"                        json:"decription"`
	PackageName string         `gorm:"column:package_name"                      json:"-"`
	CreatedAt   string         `gorm:"column:created_at"                        json:"-"`
	UpdatedAt   string         `gorm:"column:updated_at"                        json:"updated_at"`
	Branch      pq.StringArray `gorm:"column:branch;type:text[];default:'{}'"   json:"-"`
}

func (c communityPkgImpl) toBasePkgDO(pkg *domain.BasePackage, do *basePkgDO) {
	*do = basePkgDO{
		Id:          uuid.New(),
		Org:         pkg.Repository.Org,
		Repo:        pkg.Repository.Repo,
		Platform:    pkg.Repository.Platform,
		Community:   pkg.Repository.Community.Community(),
		Decription:  pkg.Repository.Desc.PackageDescription(),
		PackageName: pkg.Name.PackageName(),
		CreatedAt:   utils.Date(),
		UpdatedAt:   utils.Date(),
		Branch:      toStringArray(pkg.Branches),
	}
}

func toStringArray(v []domain.BasePackageBranch) (pq pq.StringArray) {
	pq = make([]string, len(v))

	for i := range v {
		pq[i] = v[i].String()
	}

	return
}

func (b basePkgDO) toBasePkg() (v domain.BasePackage, err error) {
	v.Id = b.Id.String()
	if v.Name, err = dp.NewPackageName(b.PackageName); err != nil {
		return
	}

	v.Repository = domain.PackageRepository{
		Org:      b.Org,
		Repo:     b.Repo,
		Platform: b.Platform,
		Desc:     dp.NewDescription(b.Decription),
	}

	if v.Repository.Community, err = dp.NewCommunity(b.Community); err != nil {
		return
	}

	for i := range b.Branch {
		v.Branches = append(v.Branches, domain.StringToBasePackageBranch(b.Branch[i]))
	}

	return
}
