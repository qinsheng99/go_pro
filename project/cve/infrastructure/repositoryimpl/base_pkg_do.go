package repositoryimpl

import (
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/utils"
)

type basePkgDO struct {
	Id          uuid.UUID      `gorm:"column:uuid;type:uuid"                    json:"-"`
	Org         string         `gorm:"column:org"                               json:"-"`
	Repo        string         `gorm:"column:repo"                              json:"repo"`
	Version     string         `gorm:"column:version"                           json:"-"`
	Platform    string         `gorm:"column:platform"                          json:"-"`
	Community   string         `gorm:"column:community"                         json:"-"`
	Decription  string         `gorm:"column:decription"                        json:"decription"`
	PackageName string         `gorm:"column:package_name"                      json:"-"`
	CreatedAt   int64          `gorm:"column:created_at"                        json:"-"`
	UpdatedAt   int64          `gorm:"column:updated_at"                        json:"updated_at"`
	Branch      pq.StringArray `gorm:"column:branch;type:text[];default:'{}'"   json:"-"`
}

func (c communityPkgImpl) toBasePkgDO(pkg domain.BasePackage) []basePkgDO {
	var res = make([]basePkgDO, 0)
	var versionMap = make(map[string][]string)
	for _, p := range pkg.Branches {
		if b, ok := versionMap[p.UpstreamVersion]; !ok {
			versionMap[p.UpstreamVersion] = []string{p.Branch}
		} else {
			versionMap[p.UpstreamVersion] = append(b, p.Branch)
		}
	}

	for ver, b := range versionMap {
		do := basePkgDO{
			Id:          uuid.New(),
			Org:         pkg.Repository.Org,
			Repo:        pkg.Repository.Repo,
			Version:     ver,
			Platform:    pkg.Repository.Platform,
			Community:   pkg.Repository.Community.Community(),
			Decription:  pkg.Repository.Desc.PackageDescription(),
			PackageName: pkg.Name.PackageName(),
			CreatedAt:   utils.Now(),
			UpdatedAt:   utils.ZeroNow(),
			Branch:      b,
		}

		res = append(res, do)
	}

	return res
}
