package _1

/**
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/
func removeElements(head *ListNode, val int) *ListNode {
	//如果头结点就是需要删除的节点 那删除了头节点后 头结点的第二个节点就是头结点
	for head != nil && head.Val == val {
		deleteNode := head
		head = deleteNode.Next
		deleteNode.Next = nil
	}
	if head == nil {
		return head
	}
	pre := head

	for pre.Next != nil {
		if pre.Next.Val == val {
			deleteNode := pre.Next
			pre.Next = deleteNode.Next
			deleteNode.Next = nil
		} else {
			pre = pre.Next
		}
	}
	return head
}
