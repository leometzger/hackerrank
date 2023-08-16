package datastructures

// https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/problem?isFullScreen=false

type BSTNode struct {
	data  int32
	left  *BSTNode
	right *BSTNode
}

func findPath(root *BSTNode, value int32, currentPath []int32) []int32 {
	currentPath = append(currentPath, root.data)

	if root.data == value {
		return currentPath
	}

	if root.data < value {
		return findPath(root.right, value, currentPath)
	}

	return findPath(root.left, value, currentPath)
}

func lca(root *BSTNode, v1 int32, v2 int32) int32 {
	pathv1 := findPath(root, v1, []int32{})
	pathv2 := findPath(root, v2, []int32{})

	for i := len(pathv1) - 1; i >= 0; i-- {
		for j := len(pathv2) - 1; j >= 0; j-- {
			if pathv1[i] == pathv2[j] {
				return pathv1[i]
			}
		}
	}

	return -1
}
