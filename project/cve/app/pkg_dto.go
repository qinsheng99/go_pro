package app

import "github.com/qinsheng99/go-domain-web/project/cve/domain"

type CmdToBasePkg = domain.BasePackage
type CmdToApplicationPkg = domain.ApplicationPackage

type Package = domain.Package
type PackageRepository = domain.PackageRepository
type BasePackageBranch = domain.BasePackageBranch

type ListBasePkgsDTO struct {
	Id            string          `json:"id"`
	Community     string          `json:"community"`
	PackageName   string          `json:"package_name"`
	VersionBranch []VersionBranch `json:"version_branch"`
}

type VersionBranch struct {
	Branch  string `json:"branch"`
	Version string `json:"version"`
}

type ListApplicationPkgsDTO struct {
	Community string `json:"community"`
	Repo      []Repo `json:"repo"`
}

type Repo struct {
	Repo string `json:"repo"`
	Pkg  Pkg    `json:"pkg"`
}

type Pkg struct {
	PackageName string `json:"package_name"`
	Version     string `json:"version"`
}

func toListBasePkgsDTO(v []domain.BasePackage) ([]ListBasePkgsDTO, error) {
	var res = make([]ListBasePkgsDTO, len(v))

	for i := range v {
		res[i] = ListBasePkgsDTO{
			Id:            v[i].Id,
			Community:     v[i].Repository.Community.Community(),
			PackageName:   v[i].Name.PackageName(),
			VersionBranch: toVersionBranch(v[i].Branches),
		}
	}

	return res, nil
}

func toVersionBranch(v []domain.BasePackageBranch) (res []VersionBranch) {
	res = make([]VersionBranch, len(v))

	for i := range v {
		res[i] = VersionBranch{
			Branch:  v[i].Branch,
			Version: v[i].UpstreamVersion,
		}
	}

	return
}

func toListApplicationDTO(v []domain.ApplicationPackage) (res []ListApplicationPkgsDTO) {
	for _, p := range v {
		res = append(res, ListApplicationPkgsDTO{
			Community: p.Repository.Community.Community(),
			Repo:      toRepo(p.Repository.Repo, p.Packages),
		})
	}

	return
}

func toRepo(repo string, pkgs []domain.Package) (res []Repo) {
	res = make([]Repo, len(pkgs))

	for i := range pkgs {
		res[i] = Repo{
			Repo: repo,
			Pkg:  toPkg(pkgs[i]),
		}
	}

	return
}

func toPkg(v domain.Package) (res Pkg) {
	res.PackageName = v.Name.PackageName()
	res.Version = v.Version

	return
}
