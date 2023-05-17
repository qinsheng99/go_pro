package task

import (
	"net/http"

	"github.com/robfig/cron/v3"

	"github.com/qinsheng99/go-domain-web/common/infrastructure/postgres"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
	cverepository "github.com/qinsheng99/go-domain-web/project/cve/infrastructure/repositoryimpl"
	"github.com/qinsheng99/go-domain-web/utils"
)

type Task struct {
	cfg     Config
	cli     utils.ReqImpl
	cron    *cron.Cron
	pkgimpl repository.PkgImpl
}

const (
	application = "application"
	base        = "base"
)

func NewTask(cfg *Config, pcfg *postgres.Config) *Task {
	return &Task{
		cfg:     *cfg,
		cli:     utils.NewRequest(&http.Transport{}),
		cron:    cron.New(cron.WithSeconds()),
		pkgimpl: cverepository.NewPkgImpl(pcfg),
	}
}

func (t *Task) Register() error {
	_, err := t.cron.AddFunc(t.cfg.Pkg.Exec, t.Pkg)

	return err
}

func (t *Task) Run() {
	t.cron.Run()
}

func (t *Task) Stop() {
	t.cron.Stop()
}
