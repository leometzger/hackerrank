package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CoinChangeTestCase struct {
	n      int64
	c      []int64
	result int64
}

func TestCoinChangeProblem(t *testing.T) {
	tests := []CoinChangeTestCase{
		{
			n:      3,
			c:      []int64{8, 3, 1, 2},
			result: 3,
		},
		{
			n:      4,
			c:      []int64{1, 2, 3},
			result: 4,
		},
		{
			n:      10,
			c:      []int64{2, 5, 3, 6},
			result: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, getWays(test.n, test.c))
	}
}
