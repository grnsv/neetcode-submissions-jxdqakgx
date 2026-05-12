/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    for {
		switch {
			case p.Val < root.Val && q.Val < root.Val:
				root = root.Left
			case p.Val > root.Val && q.Val > root.Val:
				root = root.Right
			default: return root
		}
	}
}
