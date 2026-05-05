/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	l1, l2 := lists[0], lists[1]
	if len(lists) > 2 {
		l1, l2 = mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:])
	}

	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}

	return dummy.Next
}
