/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = carry
		if l1 != nil {
			curr.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curr.Val += l2.Val
			l2 = l2.Next
		}

		if curr.Val > 9 {
			curr.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
	}

	return dummy.Next
}
