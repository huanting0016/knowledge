package main

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：
输入：head = []
输出：[]
示例 3：
输入：head = [1]
输出：[1]
*/
type LinkNode struct {
	val  int
	next *LinkNode
}

func main() {

}

func change(a *LinkNode) *LinkNode {

	if a == nil || a.next == nil {
		return a
	}
	head := a
	for head.next != nil && head != nil {
		tmp := head.next
		head.next = tmp.next
		tmp.next = head
		head = head.next
	}
	return a
}
