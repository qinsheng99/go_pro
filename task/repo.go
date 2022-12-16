package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/sirupsen/logrus"
)

type Repo struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name"`
	Path     string `json:"path"`
}

func RepoTask() {
	page := 1
	for {
		url := fmt.Sprintf("https://gitee.com/api/v5/user/repos?"+
			"access_token=70edeb9a72791f73ab6555a420fc2072&sort=full_name&page=%d&per_page=20", page)

		get, err := http.Get(url)
		if err != nil {
			return
		}

		bys, be := ioutil.ReadAll(get.Body)
		if be != nil {
			return
		}

		var res []Repo
		err = json.Unmarshal(bys, &res)
		if err != nil {
			return
		}

		if len(res) == 0 {
			break
		}

		for _, re := range res {
			var r = mysql.Repo{
				RepoId:       re.Id,
				RepoName:     re.Path,
				FullRepoName: re.FullName,
				CreateTime:   time.Now(),
				UpdateTime:   time.Now(),
			}
			if !strings.Contains(r.FullRepoName, "qinsheng") {
				continue
			}

			if r.Exist() {
				err = r.Update()
			} else {
				err = r.Insert()
			}
			if err != nil {
				logrus.Error(err)
				continue
			}
		}
		page++
	}
}
