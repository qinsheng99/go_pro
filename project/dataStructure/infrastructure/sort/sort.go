package sort

import (
	"math"

	"github.com/qinsheng99/go-domain-web/project/dataStructure/domain/sort"
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
		s.swap(i, pos, arr)
	}
}

func (s sortImpl) BubblingSort(arr []int) {
	if s.validate(arr) {
		return
	}
	l := len(arr)
	for i := l - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				s.swap(j, j+1, arr)
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
				s.swap(j, j-1, arr)
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

//
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

// ShellSort 希尔排序
func (s sortImpl) ShellSort(arr []int) {
	var h = 1
	for h <= len(arr)/3 {
		h = h*3 + 1
	}
	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < len(arr); i++ {
			for j := i; j > gap-1; j -= gap {
				if arr[j] < arr[j-gap] {
					s.swap(j, j-gap, arr)
				}
			}
		}
	}
}

func (s sortImpl) MergeSort(arr []int, left, right int) {
	if left == right {
		return
	}
	//分成两部分，左右排序，左右merge
	var mid = left + (right-left)/2

	s.MergeSort(arr, left, mid)
	s.MergeSort(arr, mid+1, right)

	merge(arr, left, mid, right)
}

func merge(arr []int, l, m, r int) {
	var newArr = make([]int, r-l+1)

	var i, j, k = l, m + 1, 0

	for i <= m && j <= r {
		if arr[i] <= arr[j] {
			newArr[k] = arr[i]
			k++
			i++
		} else {
			newArr[k] = arr[j]
			k++
			j++
		}
	}

	for i <= m {
		newArr[k] = arr[i]
		k++
		i++
	}

	for j <= r {
		newArr[k] = arr[j]
		k++
		j++
	}
	for ii := 0; ii < len(newArr); ii++ {
		arr[l+ii] = newArr[ii]
	}
}

func (s sortImpl) CountSort(arr []int) []int {
	var count = make([]int, 10)
	var result = make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}

	//for i, j := 0, 0; i < len(count); i++ {
	//	for ; count[i] > 0; count[i]-- {
	//		result[j] = i
	//		j++
	//	}
	//}

	for i := len(arr) - 1; i >= 0; i-- {
		count[arr[i]] -= 1
		result[count[arr[i]]] = arr[i]
	}

	return result
}

func (s sortImpl) RadixSort(arr []int) []int {
	var count = make([]int, 10)
	var mirror = make([]int, 10)

	var cArr = make([]int, len(arr))
	copy(cArr, arr)

	for i := 0; i < findMax(getSliceMax(arr)); i++ {
		pow := int(math.Pow(10, float64(i)))
		for j := 0; j < len(arr); j++ {
			num := arr[j] / pow % 10
			count[num]++
		}
		for m := 1; m < len(count); m++ {
			count[m] = count[m] + count[m-1]
		}

		for nx := len(arr) - 1; nx >= 0; nx-- {
			num := arr[nx] / pow % 10
			count[num] -= 1
			cArr[count[num]] = arr[nx]
		}
		copy(arr, cArr)
		copy(count, mirror)
	}

	return cArr
}

func (s sortImpl) HeapSort(arr []int) []int {
	var heapSize = len(arr)
	//3, 5, 8, 1, 4, 7, 0
	//for i := 0; i < len(arr); i++ {
	//	heapInsert(arr, i)
	//}

	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, heapSize)
	}

	for heapSize > 0 {
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		heapSize--
		heapify(arr, 0, heapSize)
	}

	return arr
}
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index, heapSize int) {
	var left = index*2 + 1
	for left < heapSize {
		var largest = 0
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[largest] > arr[index] {
		} else {
			largest = index
		}
		if index == largest {
			break
		}
		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
		left = index*2 + 1
	}
}

func findMax(num int) (length int) {
	if num == 0 {
		return 1
	}
	for i := num; int(i) != 0; i /= 10 {
		length++
	}
	return
}

func getSliceMax(arr []int) (max int) {
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return
}

func (s sortImpl) swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (s sortImpl) validate(arr []int) (flag bool) {
	if len(arr) == 1 || len(arr) == 0 {
		flag = true
	}

	return
}
