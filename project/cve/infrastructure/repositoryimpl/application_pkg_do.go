package repositoryimpl

import (
	"github.com/google/uuid"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/utils"
)

type applicationPkgDO struct {
	Id          uuid.UUID `gorm:"column:uuid;type:uuid"    json:"-"`
	Org         string    `gorm:"column:org"               json:"-"`
	Repo        string    `gorm:"column:repo"              json:"repo"`
	Assignee    string    `gorm:"column:assignee"          json:"assigne"`
	Version     string    `gorm:"column:version"           json:"-"`
	Platform    string    `gorm:"column:platform"          json:"-"`
	Community   string    `gorm:"column:community"         json:"-"`
	Milestone   string    `gorm:"column:milestone"         json:"milestone"`
	Decription  string    `gorm:"column:decription"        json:"decription"`
	PackageName string    `gorm:"column:package_name"      json:"-"`
	CreatedAt   int64     `gorm:"column:created_at"        json:"-"`
	UpdatedAt   int64     `gorm:"column:updated_at"        json:"updated_at"`
}

func (c communityPkgImpl) toApplicationPkgDO(pkg domain.ApplicationPackage) []applicationPkgDO {
	var res = make([]applicationPkgDO, 0)
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
			CreatedAt:   utils.Now(),
			UpdatedAt:   utils.ZeroNow(),
		}
		if p.Assignee != nil {
			do.Assignee = p.Assignee.Account()
		}

		res = append(res, do)
	}

	return res
}
