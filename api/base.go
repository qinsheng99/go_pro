package api

type IssueOptions struct {
	Token         string `json:"access_token"`
	Repo          string `json:"repo"`
	Title         string `json:"title"`
	IssueType     string `json:"issue_type,omitempty"`
	Body          string `json:"body"`
	Assignee      string `json:"assignee,omitempty"`
	Labels        string `json:"labels,omitempty"`
	SecurityHole  bool   `json:"security_hole"`
	Collaborators string `json:"collaborators,omitempty"`
	Program       string `json:"program,omitempty"`
	Milestone     int64  `json:"milestone,omitempty"`
}

type CreateIssueReq struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Email string `json:"email"`
	Code  string `json:"code"`
}
