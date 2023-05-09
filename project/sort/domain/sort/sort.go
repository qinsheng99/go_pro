package sort

type Sort interface {
	SelectSort(arr []int)
	BubblingSort(arr []int)
	InsertSort(arr []int)
	QuickSort(arr []int, left, right int)
}
