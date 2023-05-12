package dp

import "errors"

type branch string

type Branch interface {
	Branch() string
}

func NewBranch(v string) (Branch, error) {
	if len(v) == 0 {
		return nil, errors.New("empty branch")
	}

	return branch(v), nil
}

func (b branch) Branch() string {
	return string(b)
}
