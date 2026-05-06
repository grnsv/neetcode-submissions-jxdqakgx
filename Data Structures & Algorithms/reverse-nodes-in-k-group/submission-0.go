/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	for range k {
		if curr == nil {
			return head
		}
		curr = curr.Next
	}

	curr = reverseKGroup(curr, k)

	prev := curr
	curr = head
	for range k {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
