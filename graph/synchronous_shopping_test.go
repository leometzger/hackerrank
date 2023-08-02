package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SynchronousShoppingTestCase struct {
	n       int32
	k       int32
	centers []string
	roads   [][]int32
	result  int32
}

func TestSynchronousShopping(t *testing.T) {
	tests := []SynchronousShoppingTestCase{
		{
			n:      5,
			k:      3,
			result: 20,
			centers: []string{
				"1 1",
				"2 1 2",
				"2 2 3",
				"1 2",
				"1 2",
			},
			roads: [][]int32{
				{1, 2, 10},
				{1, 3, 15},
				{1, 4, 1},
				{2, 3, 10},
				{3, 5, 5},
			},
		},
		{
			n:      5,
			k:      5,
			result: 30,
			centers: []string{
				"1 1",
				"1 2",
				"1 3",
				"1 4",
				"1 5",
			},
			roads: [][]int32{
				{1, 2, 10},
				{1, 3, 10},
				{2, 4, 10},
				{3, 5, 10},
				{4, 5, 10},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, shop(test.n, test.k, test.centers, test.roads), test.result)
	}
}
