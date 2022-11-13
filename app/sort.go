package app

import "github.com/qinsheng99/go-domain-web/domain/sort"

type sortService struct {
	s sort.Sort
}

type SortServiceImpl interface {
	Select(arr []int)
	Bubbling(arr []int)
	Insert(arr []int)
	Quick(arr []int)
}

func NewSortService(s sort.Sort) SortServiceImpl {
	return sortService{s: s}
}

func (s sortService) Select(arr []int) {
	s.s.SelectSort(arr)
}

func (s sortService) Bubbling(arr []int) {
	s.s.BubblingSort(arr)
}

func (s sortService) Insert(arr []int) {
	s.s.InsertSort(arr)
}

func (s sortService) Quick(arr []int) {
	s.s.QuickSort(arr, 0, len(arr)-1)
}
