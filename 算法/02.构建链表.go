package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func main() {
	nums := []int{1, 2, 3}
	res := &ListNode{
		Value: nums[0],
	}
	tmp := res
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Value: nums[i]}
		tmp = tmp.Next
	}
	fmt.Println(res)
}
