package controller

type base struct {
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
}

func (base) Response(data interface{}, page, size, total int) base {
	return base{
		Data:    data,
		Total:   total,
		Page:    page,
		PerPage: size,
	}
}
