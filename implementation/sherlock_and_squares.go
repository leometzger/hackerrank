package implementation

import "math"

// https://www.hackerrank.com/challenges/sherlock-and-squares/problem?isFullScreen=false

func squares(a int32, b int32) int32 {
	var counter int32 = 0
	var first int32 = 0

	for i := a; i <= b; i++ {
		sqrt := math.Sqrt(float64(i))

		if math.Ceil(sqrt) == math.Floor(sqrt) {
			first = int32(sqrt)
			break
		}
	}

	if first != 0 {
		for i := first; i*i <= b; i++ {
			counter++
		}
	}

	return counter
}
