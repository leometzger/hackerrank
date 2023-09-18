package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ServiceLaneTestCase struct {
	width  []int32
	cases  [][]int32
	result []int32
}

func TestServiceLane(t *testing.T) {
	tests := []ServiceLaneTestCase{
		{
			width:  []int32{2, 3, 1, 2, 3, 2, 3, 3},
			cases:  [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}},
			result: []int32{1, 2, 3, 2, 1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, serviceLane(test.width, test.cases))
	}
}
