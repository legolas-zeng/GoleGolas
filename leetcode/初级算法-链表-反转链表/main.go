package main

import "fmt"

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 声明一个对象prev，此时值为nil
	var prev *ListNode
	for head != nil {
		fmt.Println(head, prev)
		//第一次循环赋值前，head=1，prev=nil
		head.Next, head, prev = prev, head.Next, head
		fmt.Println("--------")
		fmt.Println(head, prev)
		//第一次循环赋值后，head=2，prev=1
	}
	return prev
}

/////////////////////  双指针  /////////////////////
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//声明一个newhead，此时的newhead为nil
	var newHead *ListNode
	for head != nil {
		fmt.Println("交换前：", head, newHead)
		//备份head.Next给node
		node := head.Next
		fmt.Println("交换中：", node)
		//让head指向newHead
		head.Next = newHead
		//把head变成newHead
		newHead = head
		//让老的head，获取到备份的那个next
		head = node
		fmt.Println("交换后：", head, newHead, node)
	}

	return newHead
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reverseList2(head)
	fmt.Println(pre)
}
