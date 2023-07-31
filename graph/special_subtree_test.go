package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpecialSubtreeTestCase struct {
	n      int32
	start  int32
	edges  [][]int32
	result int32
}

func TestSpecialSubtree(t *testing.T) {
	tests := []SpecialSubtreeTestCase{
		{
			n:      5,
			start:  1,
			result: 15,
			edges: [][]int32{
				{1, 2, 3},
				{1, 3, 4},
				{4, 2, 6},
				{5, 2, 2},
				{2, 3, 5},
				{3, 5, 7},
			},
		},
		{
			n:      4,
			start:  4,
			result: 200,
			edges: [][]int32{
				{1, 2, 1},
				{3, 2, 150},
				{4, 3, 99},
				{1, 4, 100},
				{3, 1, 200},
			},
		},
		{
			n:      4,
			start:  2,
			result: 1106,
			edges: [][]int32{
				{2, 1, 1000},
				{3, 4, 299},
				{2, 4, 200},
				{2, 4, 100},
				{3, 2, 300},
				{3, 2, 6},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, prims(test.n, test.edges, test.start))
	}
}
