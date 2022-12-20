package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-domain-web/api"
	"github.com/qinsheng99/go-domain-web/infrastructure/mysql"
	"github.com/qinsheng99/go-domain-web/utils"
)

type base struct {
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
}

var Base = base{}

func (base) Response(data interface{}, page, size, total int) base {
	return base{
		Data:    data,
		Total:   total,
		Page:    page,
		PerPage: size,
	}
}

var emailReg = regexp.MustCompile("\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*")

func (base) SendEmail(c *gin.Context) {
	if !emailReg.MatchString(c.Param("email")) {
		utils.Failure(c, fmt.Errorf("incorrect email format,email:%s", c.Param("email")))
		return
	}

	code := utils.GenerateCode(6)
	var email = mysql.Email{
		Email:    c.Param("email"),
		IsDelete: 0,
	}

	var err error
	if email.Exist() {
		if time.Now().Unix()-email.CreateTime.Unix() < 60 {
			utils.Failure(c, fmt.Errorf("一分钟内请勿多次触发"))
			return
		}
		email.Code = code
		email.CreateTime = time.Now()
		err = email.Update()
	} else {
		email.Code = code
		email.CreateTime = time.Now()
		err = email.Insert()
	}

	if err != nil {
		utils.Failure(c, err)
		return
	}

	utils.Success(c, http.StatusOK, "")
}

func (base) VerifyCode(c *gin.Context) {
	email := mysql.Email{
		Email:    c.Param("email"),
		Code:     c.Param("code"),
		IsDelete: 0,
	}

	if !email.Check() {
		utils.Failure(c, fmt.Errorf("email code verify failed"))
		return
	}

	email.DeleteCode()
	utils.Success(c, http.StatusOK, "success")

}

func (base) CreateIssue(c *gin.Context) {
	var req api.CreateIssueReq
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		utils.Failure(c, err)
		return
	}
	email := mysql.Email{
		Email:    req.Email,
		Code:     req.Code,
		IsDelete: 0,
	}
	if !email.Check() {
		utils.Failure(c, fmt.Errorf("email code verify failed"))
		return
	}

	url := "https://gitee.com/api/v5/repos/qinsheng99/issues"

	var r = mysql.Repo{RepoId: req.Id}

	var option = api.IssueOptions{
		Token: "70edeb9a72791f73ab6555a420fc2072",
		Repo:  r.FindRepoName(),
		Title: req.Title,
		Body:  req.Body,
	}

	bys, err := json.Marshal(option)
	if err != nil {
		utils.Failure(c, err)
		return
	}

	h := utils.NewRequest(nil)
	var res map[string]interface{}
	_, err = h.CustomRequest(url, "POST", bys, nil, nil, false, &res)
	if err != nil {
		if err != nil {
			utils.Failure(c, err)
			return
		}
	}
	email.DeleteCode()
	fmt.Println(res["id"])
	utils.Success(c, http.StatusOK, "")
}
