/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var strs []string

	var dfs func (node *TreeNode)
	dfs = func (node *TreeNode) {
		if node == nil {
			strs = append(strs, "N")
			return
		}

		strs = append(strs, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return strings.Join(strs, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, " ")
	if len(strs) == 0 {
		return nil
	}

	i := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if strs[i] == "N" {
			i++
			return nil
		}

		val, _ := strconv.Atoi(strs[i])
		i++

		return &TreeNode{
			Val:   val,
			Left:  dfs(),
			Right: dfs(),
		}
	}

	return dfs()
}
