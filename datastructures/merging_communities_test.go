package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MergingCommunitiesTestCase struct {
	n        int32
	commands []string
	params   [][]int32
	result   []int32
}

func TestMergingCommunities(t *testing.T) {
	tests := []MergingCommunitiesTestCase{
		{
			n: 3,
			commands: []string{
				"Q",
				"M",
				"Q",
				"M",
				"Q",
				"Q",
			},
			params: [][]int32{
				{1},
				{1, 2},
				{2},
				{2, 3},
				{3},
				{2},
			},
			result: []int32{1, 2, 3, 3},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, MergingCommunities(test.n, test.commands, test.params))
	}
}
