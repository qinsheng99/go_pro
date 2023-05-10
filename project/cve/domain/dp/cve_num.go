package dp

import (
	"errors"
	"regexp"
)

type cveNum string

var cveReg = regexp.MustCompile(`^CVE-\d+-\d+$`)

type CVENum interface {
	CVENum() string
}

func NewCVENum(v string) (CVENum, error) {
	if !cveReg.MatchString(v) {
		return nil, errors.New("invalid CVE num")
	}

	return cveNum(v), nil
}

func (c cveNum) CVENum() string {
	return string(c)
}
