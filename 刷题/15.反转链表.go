package main

/*
输入一个链表，反转链表后，输出新链表的表头。
输入：{1,2,3}
返回值：{3,2,1}

*/

func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil {
		return pHead
	}
	cur := pHead
	pre := pHead.Next
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}
	pHead.Next = nil
	return cur

}
