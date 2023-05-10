package dp

import "errors"

type purl string

type Purl interface {
	Purl() string
}

func NewPurl(v string) (Purl, error) {
	if len(v) == 0 {
		return nil, errors.New("purl is empty")
	}

	return purl(v), nil
}

func (p purl) Purl() string {
	return string(p)
}
