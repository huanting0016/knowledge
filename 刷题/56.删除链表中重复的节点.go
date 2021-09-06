package main

/*
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
例如，链表1->2->3->3->4->4->5 处理后为 1->2->5

输入：
{1,2,3,3,4,4,5}
返回值：
{1,2,5}
*/

func deleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	res := &ListNode{}
	res.Next = pHead
	pre := res
	for pHead != nil && pHead.Next != nil {
		if pHead.Val == pHead.Next.Val {
			for pHead.Next != nil && pHead.Val == pHead.Next.Val {
				pHead.Next = pHead.Next.Next
			}
			pre.Next = pHead.Next
		} else {
			pre = pre.Next
		}

		pHead = pHead.Next

	}
	return res.Next
}
