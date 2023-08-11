package datastructures

import (
	"strings"
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

func TestCastleOnGridWithoutX(t *testing.T) {
	// Grid 100x100 with .
	var grid []string
	for i := 0; i < 100; i++ {
		var str []string
		for j := 0; j < 100; j++ {
			str = append(str, ".")
		}

		grid = append(grid, strings.Join(str, " "))
	}

	assert.Equal(t, int32(2), minimumMoves(grid, 0, 0, 99, 99))
}
