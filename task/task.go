package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"github.com/qinsheng99/go-domain-web/utils"
)

type Task struct {
	cfg  Config
	cli  utils.ReqImpl
	cron *cron.Cron
}

func NewTask(cfg *Config) *Task {
	cfg.Pkg.Endpoint = strings.TrimSuffix(cfg.Pkg.Endpoint, "/") + "/v1/pkg/upload/"

	return &Task{
		cfg:  *cfg,
		cli:  utils.NewRequest(&http.Transport{}),
		cron: cron.New(),
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

func (t *Task) Pkg() {
	for _, v := range t.cfg.Pkg.Packages {
		var resp = PkgResponse{Community: v.Community, Platform: v.Platform, Org: v.Org}
		for _, u := range v.Url {
			var res map[string]map[string]map[string]string
			_, err := t.cli.CustomRequest(
				u, http.MethodGet, nil, nil, nil, true, &res,
			)
			if err != nil {
				logrus.Errorf(
					"get community pkg failed, community:%s, url:%s ", v.Community, u,
				)

				continue
			}

			resp.PackageInfo = append(resp.PackageInfo, t.applicationPkgInfo(res)...)
		}

		if len(resp.PackageInfo) == 0 {
			logrus.Errorf("pkg record empty, community:%s", v.Community)
		}

		if err := t.toCveServer(resp, t.cfg.Pkg.Endpoint+v.Type); err != nil {
			logrus.Errorf("send pkg failed, err: %s", err.Error())
		}
	}
}

func (t *Task) applicationPkgInfo(res map[string]map[string]map[string]string) (data []pkgInfo) {
	for RepoKey, RepoValue := range res {
		for k, v := range RepoValue {
			data = append(data, pkgInfo{
				Repo:        RepoKey,
				Version:     v["version"],
				Assigne:     v["handler"],
				RepoDesc:    "",
				Milestone:   v["milestone"],
				PackageName: k,
				Branch:      []string{},
			})
		}
	}

	return
}

// example:
// [
//
//	"libopenraw": {
//	     "branch_detail": [
//	          {
//	               "summary": "Support digital camera RAW files",
//	               "package_name": "libopenraw",
//	               "description": [
//	                    "Libopenraw is a desktop agnostic effort "
//	               ],
//	               "version": "0.1.3",
//	               "brname": "master"
//	          }
//	     ]
//	}
//
// ]
func (t *Task) basePkgInfo(res []interface{}) (pkginfo []pkgInfo) {
	for _, pkg := range res {
		for repKey, repValue := range pkg.(map[string]interface{}) {
			var data = make(map[string]pkgInfo)
			if v, ok := repValue.(map[string]interface{})["branch_detail"]; ok {
				branchs := t.branchMap(v.([]interface{}))
				for _, i := range v.([]interface{}) {
					pkgv := i.(map[string]interface{})
					var info = pkgInfo{
						Repo:        repKey,
						PackageName: repKey,
					}
					if ver, ok := pkgv["version"].(string); ok {
						if len(ver) == 0 {
							continue
						}
						info.Branch = strings.Split(branchs[ver], ",")
						info.Version = ver
					}

					if desc, ok := pkgv["description"].([]interface{}); ok && len(desc) > 0 {
						info.RepoDesc = desc[0].(string)
					}

					if _, ok = data[info.Version]; !ok {
						data[info.Version] = info
					}
				}

				for _, info := range data {
					pkginfo = append(pkginfo, info)
				}
			}
		}
	}

	return
}

func (t *Task) branchMap(v []interface{}) map[string]string {
	var branchs = make(map[string]string)
	for _, i := range v {
		pkgv := i.(map[string]interface{})
		if ver, ok := pkgv["version"].(string); ok && len(ver) > 0 {
			if br, ok := pkgv["brname"].(string); ok && len(br) > 0 {
				if b, ok := branchs[ver]; ok {
					branchs[ver] = b + "," + br
				} else {
					branchs[ver] = br
				}
			}
		}
	}

	return branchs
}

func (t *Task) toCveServer(v PkgResponse, url string) error {
	bys, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = t.cli.CustomRequest(
		url, http.MethodPost, bys, nil, nil, true, nil,
	)

	if err != nil {
		return fmt.Errorf("community:%s, data:%x, err:%s", v.Community, bys, err.Error())
	}

	return err
}
