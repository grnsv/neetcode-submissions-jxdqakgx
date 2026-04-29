/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
	l := dummy
	r := dummy
	for range n {
		r = r.Next
	}

	for r.Next != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next

	return dummy.Next
}
