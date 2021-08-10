package main

/*
输入一个链表的头节点，按链表从尾到头的顺序返回每个节点的值（用数组返回）。

如输入{1,2,3}的链表
返回一个数组为[3,2,1]
0 <= 链表长度 <= 1000

输入：{67,0,24,58}
返回值：[58,24,0,67]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil || head.Next == nil {
		return []int{}
	}
	var res []int
	pre := head
	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
