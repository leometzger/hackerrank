package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ClimbingTheLeaderboard struct {
	ranked []int32
	player []int32
	result []int32
}

func TestClimbingTheLeaderboard(t *testing.T) {
	tests := []ClimbingTheLeaderboard{
		{
			ranked: []int32{100, 100, 50, 40, 40, 20, 10},
			player: []int32{5, 25, 50, 120},
			result: []int32{6, 4, 2, 1},
		},
		{
			ranked: []int32{100, 90, 90, 80, 75, 60},
			player: []int32{50, 65, 77, 90, 102},
			result: []int32{6, 5, 4, 2, 1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, climbingLeaderboard(test.ranked, test.player))
	}
}

func TestClimbingTheLeaderboardFiles(t *testing.T) {}
