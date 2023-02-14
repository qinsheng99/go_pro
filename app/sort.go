package app

import (
	"github.com/qinsheng99/go-domain-web/domain"
	"github.com/qinsheng99/go-domain-web/domain/sort"
)

type sortService struct {
	s sort.Sort
}

type SortServiceImpl interface {
	Select(arr domain.SortDP)
	Bubbling(arr domain.SortDP)
	Insert(arr domain.SortDP)
	Quick(arr domain.SortDP)
}

func NewSortService(s sort.Sort) SortServiceImpl {
	return sortService{s: s}
}

func (s sortService) Select(arr domain.SortDP) {
	s.s.SelectSort(arr.Fields.SortField())
}

func (s sortService) Bubbling(arr domain.SortDP) {
	s.s.BubblingSort(arr.Fields.SortField())
}

func (s sortService) Insert(arr domain.SortDP) {
	s.s.InsertSort(arr.Fields.SortField())
}

func (s sortService) Quick(arr domain.SortDP) {
	s.s.QuickSort(arr.Fields.SortField(), 0, len(arr.Fields.SortField())-1)
}
