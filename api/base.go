package api

type Pages struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type Sort struct {
	Data []int `json:"data"`
}

func (p *Pages) SetDefault() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Size <= 0 {
		p.Size = 10
	}
}
