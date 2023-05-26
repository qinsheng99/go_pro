package sort

type Sort interface {
	SelectSort(arr []int)
	BubblingSort(arr []int)
	InsertSort(arr []int)
	QuickSort(arr []int, left, right int)
	ShellSort(arr []int)
	MergeSort(arr []int, left, right int)
	CountSort(arr []int) []int
	RadixSort(arr []int) []int
	HeapSort(arr []int) []int
}
