package gitee

import (
	"time"

	_const "github.com/qinsheng99/go-domain-web/utils/const"
)

type PullRequest struct {
	Id              int           `json:"id"`
	Url             string        `json:"url"`
	HtmlUrl         string        `json:"html_url"`
	DiffUrl         string        `json:"diff_url"`
	PatchUrl        string        `json:"patch_url"`
	IssueUrl        string        `json:"issue_url"`
	CommitsUrl      string        `json:"commits_url"`
	CommentsUrl     string        `json:"comments_url"`
	Number          int           `json:"number"`
	State           string        `json:"state"`
	AssigneesNumber int           `json:"assignees_number"`
	TestersNumber   int           `json:"testers_number"`
	Assignees       []Assigness   `json:"assignees"`
	Testers         []interface{} `json:"testers"`
	Milestone       interface{}   `json:"milestone"`
	Labels          []Labels      `json:"labels"`
	Locked          bool          `json:"locked"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	ClosedAt        time.Time     `json:"closed_at"`
	User            struct {
		Id           int    `json:"id"`
		Login        string `json:"login"`
		Name         string `json:"name"`
		AvatarUrl    string `json:"avatar_url"`
		Url          string `json:"url"`
		HtmlUrl      string `json:"html_url"`
		Remark       string `json:"remark"`
		FollowersUrl string `json:"followers_url"`
		FollowingUrl string `json:"following_url"`
	} `json:"user"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Base  struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Id                int    `json:"id"`
			Login             string `json:"login"`
			Name              string `json:"name"`
			AvatarUrl         string `json:"avatar_url"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			Remark            string `json:"remark"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
		} `json:"user"`
	} `json:"base"`
}

type Labels struct {
	Id           int       `json:"id"`
	Color        string    `json:"color"`
	Name         string    `json:"name"`
	RepositoryId int       `json:"repository_id"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Assigness struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Url       string `json:"url"`
}

func (p *PullRequest) GetHtml() string {
	return p.HtmlUrl
}

func (p *PullRequest) GetState() string {
	return p.State
}

func (p *PullRequest) GetRef() string {
	return p.Base.Ref
}

func (p *PullRequest) GetLogin() string {
	return p.User.Login
}

func (p *PullRequest) GetCreate() string {
	return p.CreatedAt.Format(_const.Format)
}

func (p *PullRequest) GetUpdate() string {
	return p.UpdatedAt.Format(_const.Format)
}

func (p *PullRequest) GetTitle() string {
	return p.Title
}

func (p *PullRequest) GetBody() string {
	return p.Body
}

func (p *PullRequest) GetLabels() []Labels {
	return p.Labels
}

func (p *PullRequest) GetLabelName() (l []string) {
	for _, label := range p.Labels {
		l = append(l, label.Name)
	}
	return
}

func (p *PullRequest) GetAssignessName() (a []string) {
	for _, ass := range p.Assignees {
		a = append(a, ass.Login)
	}
	return
}
