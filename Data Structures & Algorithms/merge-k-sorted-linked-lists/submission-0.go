/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for {
		node := &ListNode{Val: 1001}
		listIndex := 0
		for i, list := range lists {
			if list != nil && list.Val < node.Val {
				node.Val = list.Val
				listIndex = i
			}
		}

		if node.Val == 1001 {
			break
		}

		lists[listIndex] = lists[listIndex].Next
		curr.Next = node
		curr = curr.Next
	}

	return dummy.Next
}
