package repositoryimpl

import (
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/utils"
)

const (
	updatedAt  = "updated_at"
	decription = "decription"
	branch     = "branch"
)

type basePkgDO struct {
	Id          uuid.UUID      `gorm:"column:uuid;type:uuid"`
	Org         string         `gorm:"column:org"`
	Repo        string         `gorm:"column:repo"`
	Platform    string         `gorm:"column:platform"`
	Community   string         `gorm:"column:community"`
	CreatedAt   string         `gorm:"column:created_at"`
	UpdatedAt   string         `gorm:"column:updated_at"`
	Decription  string         `gorm:"column:decription"`
	PackageName string         `gorm:"column:package_name"`
	Branch      pq.StringArray `gorm:"column:branch;type:text[];default:'{}'"`
}

func (b basePkgImpl) toBasePkgDO(pkg *domain.BasePackage, do *basePkgDO) {
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

func (b basePkgDO) toUpdatedMap() map[string]interface{} {
	return map[string]interface{}{
		updatedAt:  b.UpdatedAt,
		decription: b.Decription,
		branch:     marshalStringArray(b.Branch),
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
