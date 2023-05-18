package domain

import (
	"fmt"
	"strings"

	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
)

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

func (branch *BasePackageBranch) String() string {
	return fmt.Sprintf("%s/%v", branch.Branch, branch.UpstreamVersion)
}

func StringToBasePackageBranch(s string) (r BasePackageBranch) {
	items := strings.Split(s, "/")

	r.Branch = items[0]
	r.UpstreamVersion = items[1]

	return
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
