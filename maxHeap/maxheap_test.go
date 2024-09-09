package maxheap

import "testing"

func TestBuildMaxHeap(t *testing.T) {
	heap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	buildMaxHeap(&heap, len(heap))
	t.Log(heap)
}

func TestMaxHeapify(t *testing.T) {
	heap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	maxHeapify(&heap, 0, len(heap))
	t.Log(heap)
}

func TestHeapSort(t *testing.T) {
	heap := []int{321, 2, 31, 45, 3, 6, 22, 8, 49, 10}
	heapSort(&heap, len(heap))
	t.Log(heap)
}
