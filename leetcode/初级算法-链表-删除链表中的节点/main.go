package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	preNode := dummyHead
	for head != nil {
		next := head.Next
		if head.Val == val {
			preNode.Next = next
			break
		}
		preNode = head
		head = next
	}
	return dummyHead.Next
}

func main() {
	listnode := new(ListNode)
	val := 5
	deleteNode(listnode, val)
}
