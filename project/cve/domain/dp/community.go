package dp

import "errors"

type Community interface {
	Community() string
}

func NewCommunity(v string) (Community, error) {
	if v == "" {
		return nil, errors.New("invalid package name")
	}

	return community(v), nil
}

type community string

func (v community) Community() string {
	return string(v)
}
