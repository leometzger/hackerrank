package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PickingNumbersTestCase struct {
	a      []int32
	result int32
}

func TestPickingNumbers(t *testing.T) {
	tests := []PickingNumbersTestCase{
		{
			a:      []int32{4, 6, 5, 3, 3, 1},
			result: 3,
		},
		{
			a:      []int32{1, 2, 2, 3, 1, 2},
			result: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, pickingNumbers(test.a))
	}
}
