package sort

import (
	"github.com/qinsheng99/go-domain-web/project/sort/domain/sort"
)

type sortImpl struct {
}

func NewSort() sort.Sort {
	return sortImpl{}
}

func (s sortImpl) SelectSort(arr []int) {
	if s.validate(arr) {
		return
	}
	l := len(arr)
	for i := 0; i < l; i++ {
		var pos = i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[pos] {
				pos = j
			}
		}
		arr[i], arr[pos] = arr[pos], arr[i]
	}
}

func (s sortImpl) validate(arr []int) (flag bool) {
	if len(arr) == 1 || len(arr) == 0 {
		flag = true
	}
	return
}

func (s sortImpl) BubblingSort(arr []int) {
	if s.validate(arr) {
		return
	}
	l := len(arr)
	for i := l - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func (s sortImpl) InsertSort(arr []int) {
	if s.validate(arr) {
		return
	}
	//213
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func (s sortImpl) QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := s.partition(arr, left, right)
	s.QuickSort(arr, left, mid-1)
	s.QuickSort(arr, mid+1, right)
}

func (s sortImpl) partition(arr []int, leftBound, rightBound int) int {
	var pivot = arr[rightBound]
	var left, right = leftBound, rightBound - 1
	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[left], arr[rightBound] = arr[rightBound], arr[left]
	return left
}
