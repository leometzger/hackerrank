package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BinarySearchTreeLCATestCase struct {
	root   *BSTNode
	v1     int32
	v2     int32
	result int32
}

func TestBinarySearchTreeLCA(t *testing.T) {
	root := &BSTNode{
		data: 4,
		left: &BSTNode{
			data:  2,
			left:  &BSTNode{data: 1},
			right: &BSTNode{data: 3},
		},
		right: &BSTNode{
			data: 7,
			left: &BSTNode{data: 7},
		},
	}

	tests := []BinarySearchTreeLCATestCase{
		{
			root:   root,
			v1:     1,
			v2:     3,
			result: 2,
		},
		{
			root:   root,
			v1:     1,
			v2:     7,
			result: 4,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, lca(test.root, test.v1, test.v2))
	}
}
