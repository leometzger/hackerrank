package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.hackerrank.com/challenges/find-the-running-median/problem?isFullScreen=true

type FindRunningMedianTestCase struct {
	a      []int32
	result []float64
}

func TestFindRunningMedian(t *testing.T) {
	tests := []FindRunningMedianTestCase{
		{
			a:      []int32{7, 3, 5, 2},
			result: []float64{7.0, 5.0, 5.0, 4.0},
		},
		{
			a:      []int32{1, 2, 3, 4, 5},
			result: []float64{1.0, 1.5, 2, 2.5, 3},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, runningMedian(test.a))
	}
}
