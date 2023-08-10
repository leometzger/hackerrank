package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CastleOnTheGridTestCase struct {
	grid   []string
	startX int32
	startY int32
	goalX  int32
	goalY  int32
	result int32
}

func TestCastleOnTheGrid(t *testing.T) {
	tests := []CastleOnTheGridTestCase{
		{
			grid: []string{
				".X.",
				".X.",
				"...",
			},
			startX: 0,
			startY: 0,
			goalX:  0,
			goalY:  2,
			result: 3,
		},
		{
			grid: []string{
				"...",
				".X.",
				".X.",
			},
			startX: 2,
			startY: 0,
			goalX:  0,
			goalY:  2,
			result: 2,
		},
		{
			grid: []string{
				"...",
				".X.",
				".X.",
			},
			startX: 2,
			startY: 0,
			goalX:  2,
			goalY:  2,
			result: 3,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, minimumMoves(test.grid, test.startX, test.startY, test.goalX, test.goalY))
	}
}
