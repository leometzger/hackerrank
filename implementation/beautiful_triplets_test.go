package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type beautifulTripletsTestCase struct {
	d      int32
	arr    []int32
	result int32
}

func TestBeaultifulTriplets(t *testing.T) {
	tests := []beautifulTripletsTestCase{
		{
			result: 3,
			d:      3,
			arr:    []int32{1, 2, 4, 5, 7, 8, 10},
		},
		{
			result: 2,
			d:      3,
			arr:    []int32{1, 6, 7, 7, 8, 10, 12, 13, 14, 19},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, beautifulTriplets(test.d, test.arr))
	}
}
