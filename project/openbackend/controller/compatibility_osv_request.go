package controller

import (
	"github.com/qinsheng99/go-domain-web/domain/dp"
	"github.com/qinsheng99/go-domain-web/project/openbackend/domain"
)

type Pages struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type requestOsv struct {
	KeyWord string `json:"keyword"`
	OsvName string `json:"osvName"`
	Type    string `json:"type"`
	Pages   Pages  `json:"pages"`
}

func (p *Pages) SetDefault() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Size <= 0 {
		p.Size = 10
	}
}

func (r requestOsv) tocmd() (o domain.OsvOptions) {
	o.Page = dp.NewPage(r.Pages.Page)
	o.Size = dp.NewSize(r.Pages.Size)

	o.KeyWord = r.KeyWord
	o.OsvName = r.OsvName
	o.Type = r.Type

	return
}
