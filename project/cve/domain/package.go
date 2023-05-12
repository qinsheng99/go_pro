package domain

import "github.com/qinsheng99/go-domain-web/project/cve/domain/dp"

type PackageRepository struct {
	Org       string
	Repo      string
	Program   string
	Platform  string
	Milestone string

	Assigne  dp.Account
	RepoDesc dp.Description
}

type BasePackageBranch struct {
	Branch  dp.Branch
	Version string
}

type BasePackage struct {
	Id         string
	PkgName    dp.PackageName
	Branches   []BasePackageBranch
	Community  dp.Community
	Repository PackageRepository
}

type ApplicationPackage struct {
	Community    dp.Community
	CommunityPkg []CommunityPkg
}

type CommunityPkg struct {
	Packages   []Package
	Repository PackageRepository
}

type Package struct {
	PkgName dp.PackageName
	Version string
}
