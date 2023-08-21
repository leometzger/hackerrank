package dp

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fibonacciModifiedTestCase struct {
	t1     int8
	t2     int8
	n      int8
	result big.Int
}

func TestFibonacciModified(t *testing.T) {
	tests := []fibonacciModifiedTestCase{
		{
			t1:     0,
			t2:     1,
			n:      2,
			result: *big.NewInt(1),
		},
		{
			t1:     0,
			t2:     1,
			n:      5,
			result: *big.NewInt(5),
		},
		{
			t1:     0,
			t2:     1,
			n:      6,
			result: *big.NewInt(27),
		},
		{
			t1:     0,
			t2:     1,
			n:      6,
			result: *big.NewInt(27),
		},
		{
			t1:     0,
			t2:     1,
			n:      7,
			result: *big.NewInt(734),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, FibonacciModified(test.t1, test.t2, test.n))
	}
}
