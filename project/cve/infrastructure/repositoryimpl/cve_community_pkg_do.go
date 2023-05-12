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

type cveCommunityPkgDO struct {
	Id          uuid.UUID      `gorm:"column:uuid;type:uuid"`
	PackageName string         `gorm:"column:package_name"`
	Repo        string         `gorm:"column:repo"`
	Community   string         `gorm:"column:community" `
	Status      string         `gorm:"column:status"`
	Decription  string         `gorm:"column:decription"`
	Milestone   string         `gorm:"column:milestone"`
	Assigne     string         `gorm:"column:assigne"`
	Platform    string         `gorm:"column:platform"`
	Version     string         `gorm:"column:version"`
	CreatedAt   int64          `gorm:"column:created_at"`
	UpdatedAt   int64          `gorm:"column:updated_at"`
	Branch      pq.StringArray `gorm:"column:branch;type:text[];default:'{}'"`
}

func (c communityPkg) toAppPkgDO(v *domain.ApplicationPackage) []cveCommunityPkgDO {
	var res = make([]cveCommunityPkgDO, 0)
	for _, pkg := range v.CommunityPkg {
		for _, p := range pkg.Packages {
			res = append(res, cveCommunityPkgDO{
				Id:          uuid.New(),
				PackageName: p.PkgName.PackageName(),
				Repo:        pkg.Repository.Repo,
				Community:   v.Community.Community(),
				Status:      "add",
				Decription:  pkg.Repository.RepoDesc.Description(),
				Milestone:   pkg.Repository.Milestone,
				Assigne:     pkg.Repository.Assigne.Account(),
				Platform:    pkg.Repository.Platform,
				Version:     p.Version,
				CreatedAt:   utils.Now(),
				UpdatedAt:   utils.Now(),
				Branch:      pq.StringArray{},
			})
		}
	}

	return res
}
