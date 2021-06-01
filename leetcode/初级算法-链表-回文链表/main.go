package main

type ListNode struct {
	Val  int
	Next *ListNode
}

///////////////////// 使用转换成切片的方式 ///////////////////////////////
func isPalindrome(head *ListNode) bool {
	vallist := []int{}
	for head != nil {
		//把链表转换成切片
		vallist = append(vallist, head.Val)
		head = head.Next
	}
	n := len(vallist)
	//使用双指针法判断是否为回文
	for i, v := range vallist[:n/2] {
		if v != vallist[n-1-i] {
			return false
		}
	}
	return true
}

/////////////////////  使用快慢指针 /////////////////////////////////////

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	prev := &ListNode{
		Val:  nil,
		Next: nil,
	}
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	//把slow慢指针断开
	prev.Next = nil

	// 翻转
	var head2 *ListNode = nil

	for slow != nil {
		tmp := slow.Next
		slow.Next = head2
		head2 = slow
		slow = tmp
	}
	// 对比
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head.Next
	}

	return true
}
