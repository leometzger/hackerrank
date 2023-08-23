package dp

import "math"

// https://www.hackerrank.com/challenges/equal/problem

func equal(arr []int32) int32 {
	var ops int32

	var min int32 = math.MaxInt32
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	memo := make(map[int32]int32)
	for i := 0; i < len(arr); i++ {
		diff := arr[i] - min
		value, ok := memo[diff]
		if ok {
			ops += value
			break
		}

		ops += (diff / 5) + (diff%5)/2 + (diff%5)%2
	}

	return ops
}
