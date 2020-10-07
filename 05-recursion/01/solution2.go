package _1

/**
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	pre := dummyHead
	for pre.Next != nil {
		if pre.Next.Val == val {
			deleteNode := pre.Next
			pre.Next = deleteNode.Next
			deleteNode.Next = nil
		} else {
			pre = pre.Next
		}
	}
	return dummyHead.Next
}
