package dp

import "errors"

type sortFields []int

type SortField interface {
	SortField() []int
	Len() int
}

func (s sortFields) SortField() []int {
	return s
}

func NewSortField(arr []int) (SortField, error) {
	if len(arr) == 0 {
		return nil, errors.New("dataStructure field is empty")
	}

	return sortFields(arr), nil
}

func (s sortFields) Len() int {
	return len(s)
}
