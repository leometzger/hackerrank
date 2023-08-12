package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type QueueUsingTwoStacksTestCase struct {
	operations []string
	result     []int32
}

func TestQueueUsingTwoStacks(t *testing.T) {
	tests := []QueueUsingTwoStacksTestCase{
		{
			result: []int32{14, 14},
			operations: []string{
				"1 42",
				"2",
				"1 14",
				"3",
				"1 28",
				"3",
				"1 60",
				"1 78",
				"2",
				"2",
			},
		},
		{
			result: []int32{
				33,
				33,
				33,
			},
			operations: []string{
				"1 76",
				"1 33",
				"2",
				"1 23",
				"1 97",
				"1 21",
				"3",
				"3",
				"1 74",
				"3",
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, queueUsingTwoStacks(test.operations))
	}
}
