package main

/*
输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。

输入：
{1,3,5},{2,4,6}

返回值：
{1,2,3,4,5,6}
*/

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here

	head := &ListNode{}
	tmp := head
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val >= pHead2.Val {
			tmp.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			tmp.Next = pHead1
			pHead1 = pHead1.Next
		}
		tmp = tmp.Next

	}
	if pHead1 == nil {
		tmp.Next = pHead2
	}
	if pHead2 == nil {
		tmp.Next = pHead1
	}
	return head.Next
}
