package main

import "fmt"

func main() {
	var arr = []int{23, 54, 78, 65, 90, 89}
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) []int {
	var heapSize = len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, heapSize)
	}

	//for heapSize > 0 {
	//	arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
	//	heapSize--
	//	heapify(arr, 0, heapSize)
	//}

	return arr
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

func ShellSort(arr []int) {
	var h = 1
	for h <= len(arr)/3 {
		h = h*3 + 1
	}

	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := gap; i < len(arr); i++ {
			for j := i; j > gap-1; j -= gap {
				if arr[j] < arr[j-gap] {
					swap(j, j-gap, arr)
				}
			}
		}
	}
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, leftBound, rightBound int) int {
	var pivot = arr[rightBound]                 // 89
	var left, right = leftBound, rightBound - 1 //0,4
	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++ //4
		}
		for left <= right && arr[right] > pivot {
			right-- //3
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[left], arr[rightBound] = arr[rightBound], arr[left]

	return left
}
