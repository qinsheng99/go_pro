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
var r repository.PkgImpl

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

	r = repositoryimpl.NewPkgImpl(cfg.Postgres)

	m.Run()
}

func TestAddApplication(t *testing.T) {
	var p = domain.Package{
		Id:        "",
		Version:   "1.2.3",
		Milestone: "MT",
	}
	p.Name, _ = dp.NewPackageName("python3")
	p.Assignee, _ = dp.NewAccount("zjm")
	var a = domain.ApplicationPackage{
		Packages: []domain.Package{
			p,
		},
		Repository: domain.PackageRepository{
			Org:      "mindspore",
			Repo:     "mindspore",
			Platform: "gitee",
			Desc:     dp.NewDescription(""),
		},
	}

	a.Repository.Community, _ = dp.NewCommunity("mindspore")

	t.Log(r.AddApplicationPkg(&a))
}

func TestAddBasePkg(t *testing.T) {
	var pkg = "kernel"
	var version = "4.9.10"
	var community = "openeuler"
	var desc = fmt.Sprintf("%s security", pkg)
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
			Repo:     pkg,
			Platform: "gitee",
			Desc:     dp.NewDescription(desc),
		},
	}
	b.Repository.Community, _ = dp.NewCommunity(community)

	b.Name, _ = dp.NewPackageName(pkg)

	t.Log(r.AddBasePkg(&b))
}

func TestFindApplicationPkgs(t *testing.T) {
	c, _ := dp.NewCommunity("mindspore")

	pkg, err := r.FindApplicationPkgs(c, 0)
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

	pkg, err := r.FindApplicationPkg(opts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}

func TestFindBasePkgs(t *testing.T) {
	opt := repository.OptToFindPkgs{
		PageNum:      0,
		CountPerPage: 0,
	}

	opt.Community, _ = dp.NewCommunity("mindspore")

	pkg, err := r.FindBasePkgs(opt, 0)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}

func TestFindBasePkg(t *testing.T) {
	var (
		version   = ""
		community = ""
		name      = ""
	)

	opts := repository.OptToFindBasePkg{
		Version: version,
	}

	opts.Community, _ = dp.NewCommunity(community)
	opts.Name, _ = dp.NewPackageName(name)

	pkg, err := r.FindBasePkg(opts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v\n", pkg)
}
