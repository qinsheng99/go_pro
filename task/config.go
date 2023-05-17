package task

import (
	"github.com/qinsheng99/go-domain-web/project/cve/app"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

type Config struct {
	Pkg Package `json:"pkg"`
}

type Package struct {
	Exec     string            `json:"exec"`
	Packages []CommunityConfig `json:"packages"`
}

type CommunityConfig struct {
	Org       string   `json:"org"`
	Type      string   `json:"type"`
	Platform  string   `json:"platform"`
	Community string   `json:"community"`
	Url       []string `json:"url"`
}

type PkgResponse struct {
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

func (p *PkgResponse) toApplicationPkgCmd() (v app.CmdToApplicationPkg, err error) {
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

		v.Packages[i] = pkg
	}

	return
}

func (p *PkgResponse) toBasePkgCmd() (v []app.CmdToBasePkg, err error) {
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

		v = append(v, b)
	}

	return
}
