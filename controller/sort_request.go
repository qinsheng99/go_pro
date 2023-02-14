package controller

import (
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/dp"
)

type Sort struct {
	Data []int `json:"data"`
}

func (s Sort) tocmd() (sort domain.SortDP, err error) {
	sort.Fields, err = dp.NewSortField(s.Data)

	return
}
