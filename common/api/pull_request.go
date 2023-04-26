package api

type RequestPull struct {
	Ref       string `json:"ref"        form:"ref"`
	Sig       string `json:"sig"        form:"sig"`
	Org       string `json:"org"        form:"org"`
	Sort      string `json:"sort"       form:"sort"`
	Repo      string `json:"repo"       form:"repo"`
	Label     string `json:"label"      form:"label"`
	State     string `json:"state"      form:"state"`
	Create    string `json:"create"     form:"create"`
	Author    string `json:"author"     form:"author"`
	Search    string `json:"search"     form:"search"`
	Keyword   string `json:"keyword"    form:"keyword"`
	Assignee  string `json:"assignee"   form:"assignee"`
	Direction string `json:"direction"  form:"direction"`
	Exclusion string `json:"exclusion"  form:"exclusion"`
	Page      int    `json:"page"       form:"page"`
	PerPage   int    `json:"per_page"   form:"per_page"`
}

func (r *RequestPull) SetDefault() {
	if r.Page <= 0 {
		r.Page = 1
	}

	if r.PerPage <= 0 {
		r.PerPage = 10
	}
}
