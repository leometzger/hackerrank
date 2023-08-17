package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IsBinarySearchTreeTestCase struct {
	root   *TreeNode
	result bool
}

func TestIsBinarySearchTree(t *testing.T) {
	tests := []IsBinarySearchTreeTestCase{
		{
			result: true,
			root: &TreeNode{
				data: 10,
				left: &TreeNode{
					data:  5,
					left:  &TreeNode{data: 3},
					right: &TreeNode{data: 8},
				},
				right: &TreeNode{
					data:  20,
					left:  &TreeNode{data: 15},
					right: &TreeNode{data: 25},
				},
			},
		},
		{
			result: false,
			root: &TreeNode{
				data: 10,
				left: &TreeNode{
					data:  15,
					left:  &TreeNode{data: 23},
					right: &TreeNode{data: 8},
				},
				right: &TreeNode{
					data:  1,
					left:  &TreeNode{data: 15},
					right: &TreeNode{data: 25},
				},
			},
		},
		{
			result: false,
			root: &TreeNode{
				data: 10,
				left: &TreeNode{
					data:  5,
					left:  &TreeNode{data: 3},
					right: &TreeNode{data: 12},
				},
				right: &TreeNode{
					data:  20,
					left:  &TreeNode{data: 15},
					right: &TreeNode{data: 25},
				},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.result, checkBST(test.root))
	}
}
