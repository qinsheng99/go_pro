package elastic

type Config struct {
	Host   string `json:"host"`
	Port   int64  `json:"port"`
	Indexs indexs `json:"indexs"`
}

type indexs struct {
	PullIndex string `json:"pull_index"`
}
