package controller

import (
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/dp"
)

type RequestOsv struct {
	KeyWord string `json:"keyword"`
	OsvName string `json:"osvName"`
	Type    string `json:"type"`
	Pages   Pages  `json:"pages"`
}

func (r RequestOsv) tocmd() (o domain.OsvDP) {
	o.Page = dp.NewPage(r.Pages.Page)
	o.Size = dp.NewSize(r.Pages.Size)

	return
}
