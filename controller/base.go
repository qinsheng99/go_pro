package controller

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

type Pages struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
