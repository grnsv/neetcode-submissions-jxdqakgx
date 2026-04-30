/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := map[*Node]*Node{nil: nil}

	for curr := head; curr != nil; curr = curr.Next {
		m[curr] = &Node{Val: curr.Val}
	}

	res := m[head]
	for curr := head; curr != nil; curr = curr.Next {
		node := m[curr]
		node.Next = m[curr.Next]
		node.Random = m[curr.Random]
	}

	return res
}
