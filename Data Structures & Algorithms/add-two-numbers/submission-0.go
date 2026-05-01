/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{Val: 0}
	curr := res
	sum := l1.Val + l2.Val
	l1, l2 = l1.Next, l2.Next
	for sum > 0 {
		val := sum
		if sum >= 10 {
			val -= 10
			sum = 1
		} else {
			sum = 0
		}

		curr.Val = val

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 0 {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}

	return res
}
