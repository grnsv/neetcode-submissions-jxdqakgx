/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, res := maxDepthAndBalanced(root)
	return res
}

func maxDepthAndBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftBalanced := maxDepthAndBalanced(root.Left)
	if !leftBalanced {
		return 0, leftBalanced
	}

	rightDepth, rightBalanced := maxDepthAndBalanced(root.Right)
	if !rightBalanced {
		return 0, rightBalanced
	}

	diff := leftDepth - rightDepth
	if diff < -1 || 1 < diff {
		return 0, false
	}

	return 1 + max(leftDepth, rightDepth), true
}
