package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BetweenTwoSetsTestCase struct {
	a      []int32
	b      []int32
	result int32
}

func TestBetweenTwoSets(t *testing.T) {
	tests := []BetweenTwoSetsTestCase{
		{
			a:      []int32{2, 4},
			b:      []int32{16, 32, 96},
			result: 3,
		},
		{
			a:      []int32{2, 6},
			b:      []int32{24, 36},
			result: 2,
		},
		{
			a:      []int32{3, 4},
			b:      []int32{24, 48},
			result: 8,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, getTotalX(test.a, test.b))
	}
}
