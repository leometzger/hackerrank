package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RoadsAndLibrariesTestCase struct {
	libCost   int32
	roadCost  int32
	numCities int32
	cities    [][]int32
	result    int64
}

func TestRoadsAndLibraries(t *testing.T) {
	assert := assert.New(t)

	tests := []RoadsAndLibrariesTestCase{
		{
			numCities: 3,
			libCost:   2,
			roadCost:  1,
			cities: [][]int32{
				{1, 2},
				{3, 1},
				{2, 3},
			},
			result: 4,
		},
		{
			numCities: 6,
			libCost:   2,
			roadCost:  5,
			cities: [][]int32{
				{1, 3},
				{3, 4},
				{2, 4},
				{1, 2},
				{2, 3},
				{5, 6},
			},
			result: 12,
		},
	}

	for _, test := range tests {
		assert.Equal(test.result, RoadsAndLibraries(test.numCities, test.libCost, test.roadCost, test.cities))
	}
}
