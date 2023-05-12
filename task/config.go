package task

type Config struct {
	Pkg Package `json:"pkg"`
}

type Package struct {
	Exec     string            `json:"exec"`
	Endpoint string            `json:"endpoint"`
	Packages []CommunityConfig `json:"packages"`
}

type CommunityConfig struct {
	Org       string   `json:"org"`
	Type      string   `json:"type"`
	Platform  string   `json:"platform"`
	Community string   `json:"community"`
	Url       []string `json:"url"`
}

type PkgResponse struct {
	Org         string    `json:"org"`
	Platform    string    `json:"platform"`
	Community   string    `json:"community"`
	PackageInfo []pkgInfo `json:"package_info"`
}

type pkgInfo struct {
	Repo        string   `json:"repo"`
	Version     string   `json:"version"`
	Assigne     string   `json:"assigne"`
	RepoDesc    string   `json:"repo_desc"`
	Milestone   string   `json:"milestone"`
	PackageName string   `json:"package_name"`
	Branch      []string `json:"branch"`
}
