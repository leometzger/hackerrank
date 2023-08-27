package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type LongestCommonSubsequenceTestCase struct {
	a      []int32
	b      []int32
	result []int32
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []LongestCommonSubsequenceTestCase{
		{
			a:      []int32{1, 2, 3, 4, 1},
			b:      []int32{3, 4, 1, 2, 1, 3},
			result: []int32{3, 4, 1},
		},
		{
			a:      []int32{3, 9, 8, 3, 9, 7, 9, 7, 0},
			b:      []int32{3, 3, 9, 9, 9, 1, 7, 2, 0, 6},
			result: []int32{3, 3, 9, 9, 7, 0},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, longestCommonSubsequence(test.a, test.b))
	}
}
