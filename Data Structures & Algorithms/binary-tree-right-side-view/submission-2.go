/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    var res []int
	if root == nil {
		return res
	}

	parents := []*TreeNode{root}
	for len(parents) > 0 {
		res = append(res, parents[len(parents)-1].Val)
		children := make([]*TreeNode, 0, len(parents))
		for _, p := range parents {
			if p.Left != nil {
				children = append(children, p.Left)
			}
			if p.Right != nil {
				children = append(children, p.Right)
			}
		}

		parents = children
	}

	return res
}
