package controller

import (
	"github.com/qinsheng99/go-domain-web/project/cve/app"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
)

type pkgRequest struct {
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

type applicationPkgRequest struct {
	Data []pkgRequest `json:"data"`
}

func (a *applicationPkgRequest) toApplicationPkgCmd() (v app.CmdToApplicationPkg, err error) {
	for _, p := range a.Data {
		v.Repository.Org = p.Org
		v.Repository.Repo = p.PackageInfo[0].Repo
		v.Repository.Platform = p.Platform
		v.Repository.Desc = dp.NewDescription("")
		if v.Repository.Community, err = dp.NewCommunity(p.Community); err != nil {
			return
		}

		v.Packages = make([]app.Package, len(p.PackageInfo))

		for i := range p.PackageInfo {
			var pkg = app.Package{
				Version:   p.PackageInfo[i].Version,
				Milestone: p.PackageInfo[i].Milestone,
			}

			if pkg.Name, err = dp.NewPackageName(p.PackageInfo[i].PackageName); err != nil {
				return
			}

			if p.PackageInfo[i].Assigne != "" {
				if pkg.Assignee, err = dp.NewAccount(p.PackageInfo[i].Assigne); err != nil {
					return
				}
			}
		}
	}

	return
}

func (p *pkgRequest) toBasePkgCmd() (v app.CmdToBasePkg, err error) {
	for _, info := range p.PackageInfo {
		b := app.CmdToBasePkg{
			Repository: app.PackageRepository{
				Org:      p.Org,
				Repo:     info.Repo,
				Platform: p.Platform,
				Desc:     dp.NewDescription(info.RepoDesc),
			},
		}

		if b.Name, err = dp.NewPackageName(info.PackageName); err != nil {
			return
		}

		if b.Repository.Community, err = dp.NewCommunity(p.Community); err != nil {
			return
		}

		for _, branch := range info.Branch {
			b.Branches = append(b.Branches, app.BasePackageBranch{UpstreamVersion: info.Version, Branch: branch})
		}

		v = b
	}

	return
}

type pkgListQuery struct {
	Community    string `json:"community"       form:"community"`
	PageNum      int    `json:"page_num"        form:"page_num"`
	CountPerPage int    `json:"count_per_page"  form:"count_per_page"`
}

func (p pkgListQuery) toOptFindApplicationPkgs() (v repository.OptFindApplicationPkgs, err error) {
	if p.PageNum <= 0 {
		v.PageNum = 1
	} else {
		v.PageNum = p.PageNum
	}

	if p.CountPerPage <= 0 {
		v.CountPerPage = 100
	} else {
		v.CountPerPage = p.CountPerPage
	}

	if p.Community != "" {
		v.Community, err = dp.NewCommunity(p.Community)
	}

	return
}

func (p pkgListQuery) toOptFindBasePkgs() (v repository.OptFindBasePkgs, err error) {
	if p.PageNum <= 0 {
		v.PageNum = 1
	} else {
		v.PageNum = p.PageNum
	}

	if p.CountPerPage <= 0 {
		v.CountPerPage = 100
	} else {
		v.CountPerPage = p.CountPerPage
	}

	if p.Community != "" {
		v.Community, err = dp.NewCommunity(p.Community)
	}

	return
}
