package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReallySpecialSubtreeTestCase struct {
	gNodes  int32
	gFrom   []int32
	gTo     []int32
	gWeight []int32
	result  int32
}

func TestReallySpecialSubtree(t *testing.T) {
	tests := []ReallySpecialSubtreeTestCase{
		{
			gNodes:  4,
			gFrom:   []int32{1, 1, 4, 2, 3, 3},
			gTo:     []int32{2, 3, 1, 4, 2, 4},
			gWeight: []int32{5, 3, 6, 7, 4, 5},
			result:  12,
		},
		{
			gNodes:  5,
			gFrom:   []int32{1, 1, 1, 1, 2, 3, 4},
			gTo:     []int32{2, 3, 4, 5, 3, 4, 5},
			gWeight: []int32{20, 50, 70, 90, 30, 40, 60},
			result:  150,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, kruskals(test.gNodes, test.gFrom, test.gTo, test.gWeight))
	}
}
