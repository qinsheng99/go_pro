package dp

import (
	"errors"
	"regexp"
)

var (
	reName = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
)

type Account interface {
	Account() string
}

func NewAccount(v string) (Account, error) {
	if v == "" || !reName.MatchString(v) {
		return nil, errors.New("invalid account")
	}

	return dpAccount(v), nil
}

type dpAccount string

func (r dpAccount) Account() string {
	return string(r)
}
