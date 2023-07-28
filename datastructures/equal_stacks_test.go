package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EqualStacksTestCase struct {
	h1     []int32
	h2     []int32
	h3     []int32
	result int32
}

func TestEqualStacks(t *testing.T) {
	tests := []EqualStacksTestCase{
		{
			h1:     []int32{1, 2, 1, 1},
			h2:     []int32{1, 1, 2},
			h3:     []int32{1, 1},
			result: 2,
		},
		{
			h1:     []int32{3, 2, 1, 1, 1},
			h2:     []int32{4, 3, 2},
			h3:     []int32{1, 1, 4, 1},
			result: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, equalStacks(test.h1, test.h2, test.h3))
	}
}
