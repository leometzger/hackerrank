package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BinarySearchTreeLCATestCase struct {
	root   *TreeNode
	v1     int32
	v2     int32
	result int32
}

func TestBinarySearchTreeLCA(t *testing.T) {
	root := &TreeNode{
		data: 4,
		left: &TreeNode{
			data:  2,
			left:  &TreeNode{data: 1},
			right: &TreeNode{data: 3},
		},
		right: &TreeNode{
			data: 7,
			left: &TreeNode{data: 7},
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
