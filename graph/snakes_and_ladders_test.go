package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SnakesAndLaddersTestCase struct {
	ladders [][]int32
	snakes  [][]int32
	result  int32
}

func TestSnakesAndLadders(t *testing.T) {
	tests := []SnakesAndLaddersTestCase{
		{
			result: 3,
			ladders: [][]int32{
				{32, 62},
				{42, 68},
				{12, 98},
			},
			snakes: [][]int32{
				{95, 13},
				{97, 25},
				{93, 37},
				{79, 27},
				{75, 19},
				{49, 47},
				{67, 17},
			},
		},
		{
			result: 5,
			ladders: [][]int32{
				{8, 52},
				{6, 80},
				{26, 42},
				{2, 72},
			},
			snakes: [][]int32{
				{51, 19},
				{39, 11},
				{37, 29},
				{81, 3},
				{59, 5},
				{79, 23},
				{53, 7},
				{43, 33},
				{77, 21},
			},
		},
		{
			result:  -1,
			ladders: [][]int32{},
			snakes: [][]int32{
				{91, 90},
				{92, 90},
				{93, 90},
				{94, 90},
				{95, 90},
				{96, 90},
				{97, 90},
				{98, 90},
				{99, 90},
				{100, 90},
			},
		},
		{
			result: 3,
			ladders: [][]int32{
				{3, 54},
				{37, 100},
			},
			snakes: [][]int32{
				{56, 33},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, quickestWayUp(test.ladders, test.snakes))
	}
}
