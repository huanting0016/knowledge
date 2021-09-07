
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
