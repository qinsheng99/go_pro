package dp

import "errors"

type source string

const (
	uvp    = "uvp"
	majun  = "majun"
	vtopia = "vtopia"
)

var (
	Uvp    = source(uvp)
	Majun  = source(majun)
	Vtopia = source(vtopia)

	valiate = map[string]bool{
		uvp:    true,
		majun:  true,
		vtopia: true,
	}
)

type Source interface {
	Source() string
	IsVtopia() bool
	IsMajun() bool
	IsUvp() bool
}

func NewSource(v string) (Source, error) {
	if !valiate[v] {
		return nil, errors.New("invalid source")
	}

	return source(v), nil
}

func (s source) Source() string {
	return string(s)
}

func (s source) IsVtopia() bool {
	return s == Vtopia
}

func (s source) IsMajun() bool {
	return s == Majun
}

func (s source) IsUvp() bool {
	return s == Uvp
}
