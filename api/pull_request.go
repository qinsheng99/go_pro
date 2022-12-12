package api

type RequestPull struct {
	Page      int    `json:"page"  form:"page"`
	PerPage   int    `json:"per_page" form:"per_page"`
	Search    string `json:"search" form:"search"`
	Sort      string `json:"sort" form:"sort"`
	Sig       string `json:"sig" form:"sig"`
	State     string `json:"state" form:"state"`
	Direction string `json:"direction" form:"direction"`
	Create    string `json:"create" form:"create"`
	Author    string `json:"author" form:"author"`
	Assignee  string `json:"assignee" form:"assignee"`
	Label     string `json:"label" form:"label"`
	Exclusion string `json:"exclusion" form:"exclusion"`
	Ref       string `json:"ref" form:"ref"`
	Repo      string `json:"repo" form:"repo"`
	Org       string `json:"org" form:"org"`
	Keyword   string `json:"keyword" form:"keyword"`
	Pg        string `json:"pg" form:"pg"`
}

func (r *RequestPull) SetDefault() {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PerPage <= 0 {
		r.PerPage = 10
	}
}
