package score_api

type ScoreRes struct {
	Status  int     `json:"status"`
	Msg     string  `json:"msg"`
	Data    float64 `json:"data,omitempty"`
	Metrics metrics `json:"metrics,omitempty"`
}
type metrics struct {
	Ap   float64 `json:"ap,omitempty"`
	Ar   float64 `json:"ar,omitempty"`
	Af1  float64 `json:"af1,omitempty"`
	Af05 float64 `json:"af05,omitempty"`
	Af2  float64 `json:"af2,omitempty"`
	Acc  float64 `json:"acc,omitempty"`
	Err  float64 `json:"err,omitempty"`
}
