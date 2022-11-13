package api

type Pages struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type Sort struct {
	Data []int `json:"data"`
}
