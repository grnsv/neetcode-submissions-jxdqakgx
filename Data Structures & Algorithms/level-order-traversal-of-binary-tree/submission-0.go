/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

    var res [][]int
	parents := []*TreeNode{root}
	for len(parents) != 0 {
		s := make([]int, len(parents))
		children := make([]*TreeNode, 0, len(parents))
		
		for i, parent := range parents {
			s[i] = parent.Val
			if parent.Left != nil {
				children = append(children, parent.Left)
			}
			if parent.Right != nil {
				children = append(children, parent.Right)
			}
		}

		res = append(res, s)
		parents = children		
	}

	return res
}
