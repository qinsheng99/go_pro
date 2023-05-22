package task

import (
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/qinsheng99/go-domain-web/project/cve/app"
	"github.com/qinsheng99/go-domain-web/project/cve/domain"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/cve/domain/repository"
	"github.com/qinsheng99/go-domain-web/utils"
)

func (t *Task) ApplicationPkg() {
	for _, v := range t.cfg.Pkg.Application {
		if err := t.applicationPkg(v); err != nil {
			logrus.Errorf("application pkg failed, community:%s, err:%s", v.Community, err.Error())
		} else {
			community, _ := dp.NewCommunity(v.Community)

			if err = t.application.DeleteApplicationPkgs(repository.OptToDeleteApplicationPkgs{
				UpdatedAt: utils.Date(),
				Community: community,
			}); err != nil {
				logrus.Errorf("delete application pkg failed, community:%s, err:%s", v.Community, err.Error())
			}
		}
	}
}

func (t *Task) BasePkg() {
	for _, v := range t.cfg.Pkg.Base {
		if err := t.basePkg(v); err != nil {
			logrus.Errorf("base pkg failed, err:%s", err.Error())
		} else {
			community, _ := dp.NewCommunity(v.Community)
			if err = t.base.DeleteBasePkgs(repository.OptToDeleteApplicationPkgs{
				UpdatedAt: utils.Date(),
				Community: community,
			}); err != nil {
				logrus.Errorf("delete base pkg failed, err:%s", err.Error())
			}
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

func (t *Task) basePkg(p CommunityConfig) error {
	var resp = PkgResponse{Community: p.Community, Platform: p.Platform, Org: p.Org}
	for _, u := range p.Url {
		var res map[string]interface{}
		if _, err := t.cli.CustomRequest(
			u, http.MethodGet, nil, nil, nil, true, &res,
		); err != nil {
			return err
		}

		if res["data"] != nil {
			resp.PackageInfo = append(resp.PackageInfo, t.basePkgInfo(res["data"].([]interface{}))...)
		}
	}

	basePkg, err := resp.toBasePkgCmd()
	if err != nil {
		return err
	}

	for i := range basePkg {
		pkg := &basePkg[i]
		v, err := t.base.FindBasePkg(repository.OptToFindBasePkg{
			Community: pkg.Repository.Community,
			Name:      pkg.Name,
		})
		if err != nil {
			if err = t.base.AddBasePkg(pkg); err != nil {
				logrus.Errorf("add base pkg failed, err:%s", err.Error())
			}
		} else {
			pkg.Id = v.Id
			if err = t.base.SaveBasePkg(pkg); err != nil {
				logrus.Errorf("save base pkg failed, err:%s", err.Error())

				return err
			}
		}
	}

	return nil
}

func (t *Task) applicationPkg(p CommunityConfig) error {
	var appPkgs []app.CmdToApplicationPkg
	for _, u := range p.Url {
		bys, err := t.cli.CustomRequest(
			u, http.MethodGet, nil, nil, nil, true, nil,
		)
		if err != nil {
			return err
		}

		_ = os.Remove(p.DownloadFile)
		if err = os.WriteFile(p.DownloadFile, bys, os.ModePerm); err != nil {
			return err
		}

		var res = make(map[string]map[string]map[string]string)
		if bys, err = os.ReadFile(p.DownloadFile); err == nil {
			if p.Community == "opengauss" {
				var gauss map[string]map[string]string
				err = yaml.Unmarshal(bys, &gauss)
				res[p.Repo] = gauss
			} else {
				err = yaml.Unmarshal(bys, &res)
			}
		}

		if err != nil {
			return err
		}
		var resp = PkgResponse{
			Org:         p.Org,
			Platform:    p.Platform,
			Community:   p.Community,
			PackageInfo: t.applicationPkgInfo(res),
		}

		cmd, err := resp.toApplicationPkgCmd()
		if err != nil {
			return err
		}

		appPkgs = append(appPkgs, cmd)
	}

	for i := range appPkgs {
		for k := range appPkgs[i].Packages {
			v := appPkgs[i].Packages[k]
			repo := appPkgs[i].Repository
			appPkg, err := t.application.FindApplicationPkg(
				repository.OptToFindApplicationPkg{
					Name:      v.Name,
					Version:   v.Version,
					Repo:      repo.Repo,
					Community: repo.Community,
				},
			)
			data := domain.ApplicationPackage{
				Packages:   []domain.Package{v},
				Repository: repo,
			}
			if err == nil {
				data.Packages[0].Id = appPkg.Packages[0].Id
				if err = t.application.SaveApplicationPkg(&data); err != nil {
					logrus.Errorf("save pkg failed, err:%s\n", err.Error())

					return err
				}
			} else {
				if err = t.application.AddApplicationPkg(&data); err != nil {
					logrus.Errorf("add pkg failed, err:%s\n", err.Error())
				}
			}
		}
	}

	return nil
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
