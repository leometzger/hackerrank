package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DownToZeroIITestCase struct {
	n      int32
	result int32
}

func TestDownToZeroII(t *testing.T) {
	tests := []DownToZeroIITestCase{
		{n: 3, result: 3},
		{n: 4, result: 3},
		{n: 10, result: 5},
		{n: 266574, result: 8},
		{n: 603900, result: 7},
		{n: 847224, result: 7},
		{n: 484982, result: 9},
		{n: 937693, result: 8},
		{n: 596249, result: 9},
		{n: 991347, result: 9},
		{n: 872132, result: 9},
		{n: 386794, result: 8},
		{n: 302076, result: 9},
		{n: 51784, result: 9},
		{n: 373072, result: 9},
		{n: 580459, result: 9},
		{n: 894812, result: 9},
		{n: 42386, result: 9},
		{n: 633287, result: 10},
		{n: 309619, result: 9},
		{n: 549380, result: 8},
		{n: 667351, result: 10},
		{n: 314003, result: 9},
		{n: 855076, result: 9},
		{n: 472574, result: 11},
		{n: 332264, result: 8},
		{n: 299566, result: 8},
		{n: 344188, result: 10},
		{n: 556652, result: 8},
		{n: 226615, result: 10},
		{n: 802657, result: 9},
		{n: 5, result: 4},
		{n: 10, result: 5},
		{n: 5, result: 4},
		{n: 1, result: 1},
		{n: 1, result: 1},
		{n: 10, result: 5},
		{n: 10, result: 5},
		{n: 7, result: 5},
		{n: 8, result: 4},
		{n: 8, result: 4},
		{n: 8, result: 4},
		{n: 4, result: 3},
		{n: 1, result: 1},
		{n: 2, result: 2},
		{n: 4, result: 3},
		{n: 4, result: 3},
		{n: 9, result: 4},
		{n: 2, result: 2},
		{n: 9, result: 4},
		{n: 3, result: 3},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, downToZero(test.n))
	}
}
