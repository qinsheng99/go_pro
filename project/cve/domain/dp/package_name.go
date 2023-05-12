package dp

import (
	"errors"
)

type PackageName interface {
	PackageName() string
}

func NewPackageName(v string) (PackageName, error) {
	if v == "" {
		return nil, errors.New("invalid package name")
	}

	return packageName(v), nil
}

type packageName string

func (v packageName) PackageName() string {
	return string(v)
}
