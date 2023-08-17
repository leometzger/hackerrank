package datastructures

func checkLeft(root *TreeNode, maxAllowed int32) bool {
	if root == nil {
		return true
	}

	if root.data > maxAllowed {
		return false
	}

	return checkLeft(root.left, maxAllowed) && checkLeft(root.right, maxAllowed)
}

func checkRight(root *TreeNode, minAllowed int32) bool {
	if root == nil {
		return true
	}

	if root.data < minAllowed {
		return false
	}

	return checkRight(root.left, minAllowed) && checkRight(root.right, minAllowed)
}

func checkBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.left != nil && root.left.data > root.data {
		return false
	}

	if root.right != nil && root.right.data < root.data {
		return false
	}

	return checkLeft(root.left, root.data) &&
		checkBST(root.left) &&
		checkRight(root.right, root.data) &&
		checkBST(root.right)
}
