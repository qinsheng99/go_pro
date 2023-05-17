package repositoryimpl

import (
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/utils"
)

var pkgUpdates = []string{
	"repo", "status", "decription", "milestone", "assigne", "updated_at", "branch",
}

const (
	pkgAdd    = "add"
	pkgUpdate = "update"
	pkgDelete = "delete"
)

type communityPkgDO struct {
	Id          uuid.UUID      `gorm:"column:uuid;type:uuid"`
	Org         string         `gorm:"column:org"`
	Repo        string         `gorm:"column:repo"`
	Status      string         `gorm:"column:status"`
	Assigne     string         `gorm:"column:assigne"`
	Version     string         `gorm:"column:version"`
	Platform    string         `gorm:"column:platform"`
	Community   string         `gorm:"column:community" `
	Milestone   string         `gorm:"column:milestone"`
	Decription  string         `gorm:"column:decription"`
	PackageName string         `gorm:"column:package_name"`
	CreatedAt   int64          `gorm:"column:created_at"`
	UpdatedAt   int64          `gorm:"column:updated_at"`
	Branch      pq.StringArray `gorm:"column:branch;type:text[];default:'{}'"`
}

func (c communityPkg) toAppPkgDO(v []domain.ApplicationPackage) []communityPkgDO {
	var res = make([]communityPkgDO, 0)
	for _, pkg := range v {
		for _, p := range pkg.Packages {
			do := communityPkgDO{
				Id:          uuid.New(),
				Org:         pkg.Repository.Org,
				Repo:        pkg.Repository.Repo,
				Status:      pkgAdd,
				Version:     p.Version,
				Platform:    pkg.Repository.Platform,
				Milestone:   p.Milestone,
				Community:   pkg.Repository.Community.Community(),
				Decription:  pkg.Repository.Desc.PackageDescription(),
				PackageName: p.Name.PackageName(),
				CreatedAt:   utils.Now(),
				UpdatedAt:   utils.Now(),
				Branch:      pq.StringArray{},
			}
			if p.Assignee != nil {
				do.Assigne = p.Assignee.Account()
			}

			res = append(res, do)
		}
	}

	return res
}
