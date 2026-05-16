/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if root.Val >= right {
		return false
	}

	if root.Val <= left {
		return false
	}

	return isValid(root.Left, left, root.Val) && isValid(root.Right, root.Val, right)
}
