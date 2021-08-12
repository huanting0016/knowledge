package main

/*
输入一个链表，输出一个链表，该输出链表包含原链表中从倒数第k个结点至尾节点的全部节点。
如果该链表长度小于k，请返回一个长度为 0 的链表

输入：{1,2,3,4,5},1
返回值：{5}

*/

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	var tmp *ListNode
	pre := pHead
	cur := pHead
	for i := 0; i < k; i++ {
		if pre == nil {
			return tmp
		}
		pre = pre.Next
	}
	for pre != nil {
		cur = cur.Next
		pre = pre.Next
	}
	return cur
}
