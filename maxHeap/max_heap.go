package maxheap

func maxHeapify(heap *[]int, index int, dim int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < dim && (*heap)[left] > (*heap)[largest] {
		largest = left
	}
	if right < dim && (*heap)[right] > (*heap)[largest] {
		largest = right
	}
	if largest != index {
		(*heap)[index], (*heap)[largest] = (*heap)[largest], (*heap)[index]
		maxHeapify(heap, largest, dim)
	}
}

func buildMaxHeap(heap *[]int, dim int) {
	for i := dim / 2; i >= 0; i-- {
		maxHeapify(heap, i, dim)
	}
}

func heapSort(heap *[]int, dim int) {
	buildMaxHeap(heap, dim)
	for i := dim - 1; i > 0; i-- {
		(*heap)[0], (*heap)[i] = (*heap)[i], (*heap)[0]
		maxHeapify(heap, 0, i)
	}
}

func extractMax(heap *[]int, dim int) int {
	(*heap)[0], (*heap)[dim-1] = (*heap)[dim-1], (*heap)[0]
	maxHeapify(heap, 0, dim-1)
	return (*heap)[dim-1]
}
