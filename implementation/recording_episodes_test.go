package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RecordingEpisodesTestCase struct {
	episodes [][]int32
	result   []int32
}

func TestRecordingEpisodes(t *testing.T) {
	t.Skip()
	assert := assert.New(t)

	tests := []RecordingEpisodesTestCase{
		// // Season 1
		{
			result: []int32{1, 2},
			episodes: [][]int32{
				{10, 20, 30, 40},
				{20, 35, 21, 35},
				{14, 30, 35, 50},
			},
		},
		// Season 2
		{
			result: []int32{1, 1},
			episodes: [][]int32{
				{10, 20, 30, 40},
			},
		},
		// Season 3
		{
			result: []int32{1, 1},
			episodes: [][]int32{
				{11, 19, 31, 39},
				{12, 38, 13, 37},
				{10, 20, 30, 40},
			},
		},
		// // Season 4
		{
			result: []int32{4, 4},
			episodes: [][]int32{
				{11, 19, 31, 39},
				{12, 38, 13, 37},
				{10, 20, 30, 40},
				{10, 12, 30, 40},
				{13, 14, 30, 40},
				{15, 16, 30, 40},
				{17, 18, 30, 40},
			},
		},
		{
			result: []int32{1, 2},
			episodes: [][]int32{
				{10, 20, 30, 40},
				{20, 35, 21, 35},
				{14, 30, 35, 50},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(test.result, RecordingEpisodes(test.episodes))
	}
}
