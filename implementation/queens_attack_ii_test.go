package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type QueensAttackIITestCase struct {
	n         int32
	k         int32
	r_q       int32
	c_q       int32
	obstacles [][]int32
	result    int32
}

func TestQueensAttackII(t *testing.T) {
	tests := []QueensAttackIITestCase{
		{
			n:         1,
			k:         0,
			r_q:       1,
			c_q:       1,
			result:    0,
			obstacles: [][]int32{},
		},
		{
			n:         4,
			k:         0,
			r_q:       4,
			c_q:       4,
			result:    9,
			obstacles: [][]int32{},
		},
		{
			n:      5,
			k:      3,
			r_q:    4,
			c_q:    3,
			result: 10,
			obstacles: [][]int32{
				{5, 5},
				{4, 2},
				{2, 3},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, queensAttack(test.n, test.k, test.r_q, test.c_q, test.obstacles))
	}
}
