package implementation

func pickingNumbers(a []int32) int32 {
	index := make([]int32, 100)
	var max int32 = 0

	for i := 0; i < len(a); i++ {
		index[a[i]-1]++
	}

	for i := 1; i < len(index); i++ {
		if index[i]+index[i-1] > max {
			max = index[i] + index[i-1]
		}
	}

	return max
}
