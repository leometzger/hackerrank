package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type LongestIncresingSubsequenceTestCase struct {
	arr    []int32
	result int32
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []LongestIncresingSubsequenceTestCase{
		{
			arr:    []int32{2, 7, 4, 3, 8},
			result: 3,
		},
		{
			arr:    []int32{2, 4, 3, 7, 4, 5},
			result: 4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, longestIncreasingSubsequence(test.arr))
	}
}
