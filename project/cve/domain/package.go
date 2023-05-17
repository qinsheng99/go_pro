package domain

import "github.com/qinsheng99/go-domain-web/project/cve/domain/dp"

type PackageRepository struct {
	Org      string
	Repo     string
	Platform string

	Desc      dp.PackageDescription
	Community dp.Community
}

type BasePackageBranch struct {
	Branch          string
	UpstreamVersion string
}

type BasePackage struct {
	Id         string
	Name       dp.PackageName
	Branches   []BasePackageBranch
	Repository PackageRepository
}

type ApplicationPackage struct {
	Packages   []Package
	Repository PackageRepository
}

type Package struct {
	Id        string
	Version   string
	Milestone string

	Name     dp.PackageName
	Assignee dp.Account
}
