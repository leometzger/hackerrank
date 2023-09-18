package implementation

import "math"

// https://www.hackerrank.com/challenges/service-lane/problem?isFullScreen=false

func serviceLane(width []int32, cases [][]int32) []int32 {
	var result []int32

	for _, c := range cases {
		var min int32 = math.MaxInt32

		start := c[0]
		end := c[1]

		for i := start; i <= end; i++ {
			if width[i] < min {
				min = width[i]
			}
		}
		result = append(result, min)
	}

	return result
}
