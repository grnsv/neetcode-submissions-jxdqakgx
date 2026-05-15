/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(root *TreeNode, maxVal int) int
	dfs = func(root *TreeNode, maxVal int) int {
		if root == nil {
			return 0
		}

		res := 0
		if root.Val >= maxVal {
			res = 1
		}

		maxVal = max(maxVal, root.Val)
		return res + dfs(root.Left, maxVal) + dfs(root.Right, maxVal)
	}

	return dfs(root, root.Val)
}
