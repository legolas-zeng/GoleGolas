package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	//将下一个节点的值复制过来
	node.Val = node.Next.Val
	//将下下个节点指针复制过来
	node.Next = node.Next.Next

}

func main() {
	listnode := new(ListNode)

	deleteNode(listnode)
}
