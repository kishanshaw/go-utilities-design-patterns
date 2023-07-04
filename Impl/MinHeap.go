package Impl

import (
	"fmt"
	"log"
	"math"
)

type MinHeap struct {
	Capacity int64
	Items    []int64
}

func (h *MinHeap) Len() int64 {
	return int64(len(h.Items))
}

func leftChild(i int64) int64 {
	return 2*i + 1
}

func rightChild(i int64) int64 {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) Offer(item int64) {
	if h.Len() == h.Capacity {
		log.Println("The heap is full!")
		return
	}
	h.Items = append(h.Items, item)
	for i := len(h.Items) - 1; i != 0 && h.Items[parent(i)] > h.Items[i]; {
		h.Items[parent(i)], h.Items[i] = h.Items[i], h.Items[parent(i)]
		i = parent(i)
	}
}

func (h *MinHeap) PrintElements() {
	fmt.Print("\nPrinting Heap: ")
	for i := 0; i < len(h.Items); i++ {
		fmt.Printf(" %d", h.Items[i])
	}
	fmt.Println()
}

func (h *MinHeap) minHeapify(i int64) {
	leftChild := leftChild(i)
	rightChild := rightChild(i)
	smallest := i
	if leftChild < h.Len() && h.Items[smallest] > h.Items[leftChild] {
		smallest = leftChild
	}
	if rightChild < h.Len() && h.Items[smallest] > h.Items[rightChild] {
		smallest = rightChild
	}

	if smallest != i {
		h.Items[smallest], h.Items[i] = h.Items[i], h.Items[smallest]
		h.minHeapify(smallest)
	}
}

func (h *MinHeap) Poll() int64 {
	if h.Len() == 0 {
		return math.MaxInt64
	} else if h.Len() == 1 {
		item := h.Items[0]
		h.Items = []int64{}
		return item
	}
	item := h.Items[0]
	h.Items = h.Items[1:]
	h.minHeapify(0)
	return item
}
