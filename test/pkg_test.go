package test

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/config"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/cve/infrastructure/repositoryimpl"
)

var cfg *config.Config

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

	m.Run()
}

func TestPkg(t *testing.T) {
	r := repositoryimpl.NewPkgImpl(cfg.Postgres)

	var p = domain.Package{
		Id:        "",
		Version:   "1.2.3",
		Milestone: "MT",
	}
	p.Name, _ = dp.NewPackageName("ppp")
	p.Assignee, _ = dp.NewAccount("pppx")
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

	a.Repository.Community, _ = dp.NewCommunity("mindpsore")

	var b = domain.BasePackage{
		Name: nil,
		Branches: []domain.BasePackageBranch{
			{
				UpstreamVersion: "123",
				Branch:          "2203",
			},
			{
				UpstreamVersion: "123",
				Branch:          "2203-1",
			},
		},
		Repository: domain.PackageRepository{
			Org:       "src-openeuler",
			Repo:      "hah",
			Platform:  "gitee",
			Desc:      dp.NewDescription("oooooooooooooooooo"),
			Community: nil,
		},
	}
	b.Repository.Community, _ = dp.NewCommunity("openeuler")

	b.Name, _ = dp.NewPackageName("heihei")

	t.Log(r.AddApplicationPkg(a))
	t.Log(r.AddBasePkg(b))
}
