package controller

import (
	"github.com/qinsheng99/go-domain-web/project/cve/app"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type PkgRequest struct {
	Org         string    `json:"org"`
	Platform    string    `json:"platform"`
	Community   string    `json:"community"`
	PackageInfo []pkgInfo `json:"package_info"`
}

type pkgInfo struct {
	Repo        string   `json:"repo"`
	Version     string   `json:"version"`
	Assigne     string   `json:"assigne"`
	RepoDesc    string   `json:"repo_desc"`
	Milestone   string   `json:"milestone"`
	PackageName string   `json:"package_name"`
	Branch      []string `json:"branch"`
}

func (p *PkgRequest) toApplicationPkgCmd() (v app.CmdToApplicationPkg, err error) {
	if v.Community, err = dp.NewCommunity(p.Community); err != nil {
		return
	}

	for _, info := range p.PackageInfo {
		pkg := app.Package{Version: info.Version}
		if pkg.PkgName, err = dp.NewPackageName(info.PackageName); err != nil {
			return
		}
		if idx := v.FindRepo(info.Repo); idx != -1 {
			v.CommunityPkg[idx].Packages = append(v.CommunityPkg[idx].Packages, pkg)
		} else {
			c := app.CommunityPkg{
				Packages: []app.Package{pkg},
				Repository: domain.PackageRepository{
					Org:       p.Org,
					Repo:      info.Repo,
					Program:   "",
					Platform:  p.Platform,
					Milestone: info.Milestone,
				},
			}
			c.Repository.Assigne, _ = dp.NewAccount(info.Assigne)
			c.Repository.RepoDesc = dp.NewDescription(info.RepoDesc)
			v.CommunityPkg = append(v.CommunityPkg, c)
		}
	}

	return
}
