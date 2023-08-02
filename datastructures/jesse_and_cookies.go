package datastructures

import (
	"container/heap"
)

type CookiesHeap []int32

func (h CookiesHeap) Len() int           { return len(h) }
func (h CookiesHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h CookiesHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *CookiesHeap) Peek() int32 {
	if len(*h) == 0 {
		return -1
	}
	return (*h)[0]
}

func (h *CookiesHeap) Push(x any) {
	*h = append(*h, x.(int32))
}

func (h *CookiesHeap) Pop() any {
	old := *h
	n := len(*h)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func cookies(k int32, A []int32) int32 {
	var rounds int32 = 0
	cookieHeap := &CookiesHeap{}

	heap.Init(cookieHeap)
	for _, cookie := range A {
		heap.Push(cookieHeap, cookie)
	}

	for len(*cookieHeap) > 1 && cookieHeap.Peek() < k {
		item1 := heap.Pop(cookieHeap).(int32)
		item2 := heap.Pop(cookieHeap).(int32)
		sweetness := item1 + (2 * item2)
		heap.Push(cookieHeap, sweetness)

		rounds++
	}

	if cookieHeap.Peek() < k {
		return -1
	}

	return rounds
}
