package task

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/qinsheng99/go-domain-web/project/cve/app"
)

func (t *Task) Pkg() {
	for _, v := range t.cfg.Pkg.Packages {
		switch v.Type {
		case application:
			t.applicationPkg(v)
		case base:
			t.basePkg(v)
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

func (t *Task) basePkg(p CommunityConfig) {
	var resp = PkgResponse{Community: p.Community, Platform: p.Platform, Org: p.Org}
	var pkgs []pkgInfo
	for _, u := range p.Url {
		var res map[string]interface{}
		_, err := t.cli.CustomRequest(
			u, http.MethodGet, nil, nil, nil, true, &res,
		)
		if err != nil {
			logrus.Errorf(
				"get community pkg failed, community:%s, url:%s, err:%s ", p.Community, u, err.Error(),
			)

			continue
		}

		if res == nil || res["data"] == nil || len(res["data"].([]interface{})) == 0 {
			logrus.Errorf("empty data, url:%s", u)

			continue
		}

		fmt.Println(len(res["data"].([]interface{})))

		pkgs = append(pkgs, t.basePkgInfo(res["data"].([]interface{}))...)
	}

	if len(pkgs) == 0 {
		logrus.Errorf("pkg record empty, community:%s", p.Community)

		return
	}

	resp.PackageInfo = pkgs

	//cmd, err := resp.toBasePkgCmd()
	//if err != nil {
	//	logrus.Errorf("cmd to applicationPkg failed, err:%s", err.Error())
	//
	//	return
	//}
	//
	//if err = t.pkgimpl.AddApplicationPkg(&cmd); err != nil {
	//	logrus.Errorf("add application pkg failed, err:%s, community:%s", err, cmd[0].Community.Community())
	//}
}

func (t *Task) applicationPkg(p CommunityConfig) {
	var cmdPkg []app.CmdToApplicationPkg
	for _, u := range p.Url {
		var resp = PkgResponse{Community: p.Community, Platform: p.Platform, Org: p.Org}
		var res = make(map[string]map[string]map[string]string)
		var gauss map[string]map[string]string
		bys, err := t.cli.CustomRequest(
			u, http.MethodGet, nil, nil, nil, true, nil,
		)
		if err != nil {
			logrus.Errorf("get community pkg failed, community:%s, url:%s, err:%s", p.Community, u, err.Error())

			return
		}
		var name = "logs/pkg.yaml"
		if err = os.WriteFile(name, bys, os.ModePerm); err != nil {
			logrus.Errorf("write file failed, err:%s", err.Error())

			return
		}

		if bys, err = os.ReadFile(name); err == nil {
			if p.Community == "opengauss" {
				err = yaml.Unmarshal(bys, &gauss)

				res["security"] = gauss
			} else {
				err = yaml.Unmarshal(bys, &res)
			}
		}

		if err != nil {
			return
		}

		resp.PackageInfo = t.applicationPkgInfo(res)

		cmd, err := resp.toApplicationPkgCmd()
		if err != nil {
			logrus.Errorf("cmd to applicationPkg failed, err:%s", err.Error())

			return
		}

		cmdPkg = append(cmdPkg, cmd)

		_ = os.Remove(name)
	}

	if len(cmdPkg) == 0 {
		logrus.Errorf("pkg record empty, community:%s", p.Community)
		os.Exit(1)
		return
	}

	if err := t.pkgimpl.AddApplicationPkg(cmdPkg); err != nil {
		logrus.Errorf("add application pkg failed, err:%s, community:%s", err, p.Community)
	}
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
