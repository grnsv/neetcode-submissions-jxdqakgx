/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	_, res := maxDepthAndDiameter(root)
	return res
}

func maxDepthAndDiameter(root *TreeNode) (depth int, diameter int) {
	if root == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := maxDepthAndDiameter(root.Left)
	rightDepth, rightDiameter := maxDepthAndDiameter(root.Right)

	depth = 1 + max(leftDepth, rightDepth)
	diameter = leftDepth + rightDepth

	return depth, max(diameter, leftDiameter, rightDiameter)
}