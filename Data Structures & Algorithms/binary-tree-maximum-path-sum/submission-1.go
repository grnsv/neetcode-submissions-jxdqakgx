/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	res := root.Val

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		maxLeft := max(dfs(node.Left), 0)
		maxRight := max(dfs(node.Right), 0)

		res = max(res, node.Val + maxLeft + maxRight)

		return node.Val + max(maxLeft, maxRight)
	}

	dfs(root)

	return res
}
