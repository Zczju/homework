package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNodeByArray(head *ListNode) *ListNode {
	// 数组:
	arr := []int{}
	node := head
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	n := len(arr) / 2
	for i := 0; i < n; i++ {
		head = head.Next
	}
	return head
}

func middleNodeSinglePointer(head *ListNode) *ListNode {
	// 单指针 :
	// 第一次遍历:获取链表中节点数量(每次移动指针n++)
	// 第二次遍历:走到n/2的位置
	i := 0
	node := head
	for node != nil {
		i++
		node = node.Next
	}
	for j := 0; j < i/2; j++ {
		head = head.Next
	}

	return head
}

func middleNode2Pointer(head *ListNode) *ListNode {
	// 双指针:
	// 一快一慢，快的一次走两步，慢的一次走一步
	// 当快的走到末尾时，慢的走到1/2位置
	// 奇数情况 : 1,2,3,4,5    q从1走到3走到5，q.Next为nil，停；p从1走到2走到3
	// 偶数情况 : 1,2,3,4,5,6  q从1走到3走到5,再走一步，q为nil，停；p 从1走到2走到3走到4

	p, q := head, head              // q为快，p为慢
	for q != nil && q.Next != nil { // 循环条件要注意，容易出现空指针
		p = p.Next
		q = q.Next.Next
	}
	return p
}
