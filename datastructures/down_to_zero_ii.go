package datastructures

// https://www.hackerrank.com/challenges/down-to-zero-ii/problem

// !! Needs optimization to pass all tests !!

type DownToZeroQueue struct {
	items []*DownToZeroItem
}

type DownToZeroItem struct {
	value int32
	step  int32
}

func (q *DownToZeroQueue) isEmpty() bool                 { return len(q.items) == 0 }
func (q *DownToZeroQueue) enqueue(value *DownToZeroItem) { q.items = append(q.items, value) }
func (q *DownToZeroQueue) dequeue() *DownToZeroItem {
	if len(q.items) == 0 {
		return nil
	}
	result := q.items[0]
	q.items = q.items[1:]
	return result
}

func downToZero(n int32) int32 {
	visited := make(map[int32]bool)
	queue := &DownToZeroQueue{}
	queue.enqueue(&DownToZeroItem{
		value: 1,
		step:  1,
	})

	for !queue.isEmpty() {
		item := queue.dequeue()
		if visited[item.value] {
			continue
		}

		visited[item.value] = true

		if item.value+1 == n {
			return item.step + 1
		}

		if item.value+1 < n {
			queue.enqueue(&DownToZeroItem{
				value: item.value + 1,
				step:  item.step + 1,
			})
		}

		for i := 2; i <= int(item.value); i++ {
			if i*int(item.value) > int(n) {
				break
			}

			if item.value*int32(i) == n {
				return item.step + 1
			}

			queue.enqueue(&DownToZeroItem{
				value: item.value * int32(i),
				step:  item.step + 1,
			})
		}
	}

	return 1
}
