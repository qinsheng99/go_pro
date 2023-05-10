package dp

import "errors"

const (
	add      = "add"
	update   = "update"
	complete = "complete"
)

var (
	Add      = cveStatus(add)
	Update   = cveStatus(update)
	Complete = cveStatus(complete)

	valiateStatus = map[string]bool{
		add:      true,
		update:   true,
		complete: true,
	}
)

type cveStatus string

func NewCVEStatus(v string) (CVEStatus, error) {
	if !valiateStatus[v] {
		return nil, errors.New("invalid CVE status")
	}

	return cveStatus(v), nil
}

type CVEStatus interface {
	CVEStatus() string
}

func (c cveStatus) CVEStatus() string {
	return string(c)
}
