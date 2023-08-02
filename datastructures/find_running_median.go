package datastructures

import "container/heap"

// https://www.hackerrank.com/challenges/find-the-running-median/problem?isFullScreen=true

type FindRunningMedianHeap struct {
	items []int32
	min   bool
}

func (h FindRunningMedianHeap) Len() int      { return len(h.items) }
func (h FindRunningMedianHeap) Swap(i, j int) { h.items[i], h.items[j] = h.items[j], h.items[i] }

func (h FindRunningMedianHeap) Less(i, j int) bool {
	if h.min {
		return h.items[i] < h.items[j]
	}
	return h.items[i] > h.items[j]
}

func (h FindRunningMedianHeap) Peek() any {
	return h.items[0]
}

func (h *FindRunningMedianHeap) Push(x any) {
	h.items = append(h.items, x.(int32))
}

func (h *FindRunningMedianHeap) Pop() any {
	n := len(h.items)
	item := h.items[n-1]
	h.items = h.items[0 : n-1]
	return item
}

func runningMedian(items []int32) []float64 {
	var result []float64

	minHeap := &FindRunningMedianHeap{min: true}
	heap.Init(minHeap)

	maxHeap := &FindRunningMedianHeap{min: false}
	heap.Init(maxHeap)

	for index, i := range items {
		if (index+1)%2 == 0 {
			if i < maxHeap.Peek().(int32) {
				item := heap.Pop(maxHeap)
				heap.Push(minHeap, item)
				heap.Push(maxHeap, i)
			} else {
				heap.Push(minHeap, i)
			}

			min := minHeap.Peek().(int32)
			max := maxHeap.Peek().(int32)
			res := float64((float32(min) + float32(max)) / 2)
			result = append(result, res)
		} else {
			if len(minHeap.items) > 0 && i > minHeap.Peek().(int32) {
				item := heap.Pop(minHeap)
				heap.Push(maxHeap, item)
				heap.Push(minHeap, i)
			} else {
				heap.Push(maxHeap, i)
			}

			res := float64(maxHeap.Peek().(int32))
			result = append(result, res)
		}
	}

	return result
}
