package main

// https://www.hackerrank.com/challenges/the-hurdle-race/problem?h_r=internal-search
func HurdleRace(k int32, height []int32) int32 {
	max := int32(0)

	for _, value := range height {
		if max < value {
			max = value
		}
	}

	if k >= max {
		return 0
	}

	return max - k
}
