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
			result: []int32{1, 2, 3},
		},
	}

	for _, test := range tests {
		assert.Equal(t, int32(3), longestCommonSubsequence(test.a, test.b))
	}
}
