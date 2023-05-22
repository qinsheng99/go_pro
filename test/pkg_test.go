package test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
	"github.com/qinsheng99/go-domain-web/project/cve/infrastructure/repositoryimpl"
)

var cfg *config.Config
var base repository.BasePkgRepository
var application repository.ApplicationPkgRepository

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.Init("../config/config.yaml")
	if err != nil {
		logrus.WithError(err).Error("config init failed")

		return
	}

	if err = postgres.Init(cfg.Postgres); err != nil {
		logrus.WithError(err).Error("postgres init failed")

		return
	}

	application = repositoryimpl.NewApplicationPkgImpl(cfg.Postgres)
	base = repositoryimpl.NewBasePkgImpl(cfg.Postgres)

	m.Run()
}

func TestAddApplication(t *testing.T) {
	a := applicationPkg("mindspore", "apk", "xiaohu", "kk", "mindspore kk")

	t.Log(application.AddApplicationPkg(&a))
}

func applicationPkg(community, repo, assignee, name, desc string) domain.ApplicationPackage {
	var p = domain.Package{
		Id:        "",
		Version:   "1.2.3",
		Milestone: "MT",
	}
	p.Name, _ = dp.NewPackageName(name)
	p.Assignee, _ = dp.NewAccount(assignee)
	var a = domain.ApplicationPackage{
		Packages: []domain.Package{
			p,
		},
		Repository: domain.PackageRepository{
			Org:      community,
			Repo:     repo,
			Platform: "gitee",
			Desc:     dp.NewDescription(desc),
		},
	}

	a.Repository.Community, _ = dp.NewCommunity(community)

	return a
}

func basePkg(name string) domain.BasePackage {
	var version = "4.9.10"
	var community = "openeuler"
	var desc = fmt.Sprintf("%s security", name)
	var b = domain.BasePackage{
		Branches: []domain.BasePackageBranch{
			{
				UpstreamVersion: version,
				Branch:          "openeuler-22.03-LTS-SP1",
			},
			{
				UpstreamVersion: version,
				Branch:          "openeuler-22.03-LTS",
			},
		},
		Repository: domain.PackageRepository{
			Org:      "src-openeuler",
			Repo:     name,
			Platform: "gitee",
			Desc:     dp.NewDescription(desc),
		},
	}
	b.Repository.Community, _ = dp.NewCommunity(community)

	b.Name, _ = dp.NewPackageName(name)

	return b
}

func TestAddBasePkg(t *testing.T) {
	b := basePkg("git")
	t.Log(base.AddBasePkg(&b))
}

func TestFindApplicationPkgs(t *testing.T) {
	opt := repository.OptFindApplicationPkgs{
		PageNum:      0,
		CountPerPage: 0,
	}

	opt.Community, _ = dp.NewCommunity("mindspore")

	pkg, err := application.FindApplicationPkgs(opt)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}

func TestFindApplicationPkg(t *testing.T) {
	var repo = ""
	var version = ""
	var community = ""
	var name = ""

	opts := repository.OptToFindApplicationPkg{
		Repo:    repo,
		Version: version,
	}

	opts.Community, _ = dp.NewCommunity(community)
	opts.Name, _ = dp.NewPackageName(name)

	pkg, err := application.FindApplicationPkg(opts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}

func TestFindBasePkgs(t *testing.T) {
	opt := repository.OptFindBasePkgs{
		PageNum:      0,
		CountPerPage: 0,
	}

	opt.Community, _ = dp.NewCommunity("mindspore")

	pkg, err := base.FindBasePkgs(opt)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}

func TestFindBasePkg(t *testing.T) {
	var (
		community = ""
		name      = ""
	)

	opts := repository.OptToFindBasePkg{}

	opts.Community, _ = dp.NewCommunity(community)
	opts.Name, _ = dp.NewPackageName(name)

	pkg, err := base.FindBasePkg(opts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}

func TestSaveBasePkg(t *testing.T) {
	b := basePkg("git")
	b.Id = "9309bc5b-918c-41b4-976c-5554754150b2"

	t.Log(base.SaveBasePkg(&b))
}

func TestSaveApplicationPkg(t *testing.T) {
	a := applicationPkg(
		"mindspore", "mindspore", "xiaohu", "python3", "mindspore python3",
	)
	a.Packages[0].Id = "8ce0d85b-520d-4aac-8bed-fbfa2d02ebc5"

	t.Log(application.SaveApplicationPkg(&a))
}
