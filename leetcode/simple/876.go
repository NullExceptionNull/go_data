package simple

/**
给定一个带有头结点 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点
*/
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	m := make(map[int]*ListNode)
	var a = 0
	m[0] = head
	a++
	for head.Next != nil {
		m[a] = head.Next
		head = head.Next
		a++
	}
	var c = a / 2

	return m[c+1]
}

func middleNode2(head *ListNode) *ListNode {

	var res []*ListNode

	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	return res[len(res)/2]

}
