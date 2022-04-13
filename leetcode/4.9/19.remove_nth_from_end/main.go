package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 获取长度
	length := getLength(head)

	// 如果长度为1，只有一个节点，删掉返回nil
	if length == 1 && n > 0 {
		return nil
	}

	// 如果长度为n，删除的是第一个节点
	if length == n {
		newHead := head.Next
		head.Next = nil
		return newHead
	}

	node := head
	for i := 1; i < length-n; i++ {
		node = node.Next
	}
	deleNode := node.Next

	// 如果要删除的节点为最后一个，直接删除即可
	if deleNode.Next == nil {
		node.Next = nil
		return head
	}
	node.Next = deleNode.Next
	deleNode.Next = nil

	return head
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}
