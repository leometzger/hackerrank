package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ComponentsInGraphTestCase struct {
	gb     [][]int32
	result []int32
}

func TestComponentsInGraph(t *testing.T) {
	tests := []ComponentsInGraphTestCase{
		{
			gb:     [][]int32{{1, 5}, {1, 6}, {2, 4}},
			result: []int32{2, 3},
		},
		{
			gb:     [][]int32{{1, 6}, {2, 7}, {3, 8}, {4, 9}, {2, 6}},
			result: []int32{2, 4},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, componentsInGraph(test.gb))
	}

}
