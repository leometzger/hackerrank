package implementation

func beautifulTriplets(d int32, arr []int32) int32 {
	var counter int32 = 0
	var index map[int32]int32 = make(map[int32]int32)

	for _, n := range arr {
		index[n]++
	}

	for n, value := range index {
		diff1 := n + d
		diff2 := n + (d * 2)

		if index[diff1] > 0 && index[diff2] > 0 {
			counter += value * index[diff1] * index[diff2]
		}
	}

	return counter
}
