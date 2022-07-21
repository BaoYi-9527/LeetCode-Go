package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head} // 结果链表
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		// 当 n = 0 时， fast 正好移动到目标节点
		// 此时循环会继续进行 因为 fast 还有后续节点
		// 而后续节点数即为 链表长度 - n 正好是目标节点的正序排序
		// 也就是在循环结束时 slow 所在节点即为目标节点
		// 而 preSlow 在 slow 的上一个节点
		// 此时要删除链表中的目标节点 只需要将 preSlow 的下一个节点指向 slow 的下一个节点
		// 最后返回 dummyHead.Next
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 边界条件 链表为空
	if head == nil {
		return nil
	}
	// 边界条件
	if n <= 0 {
		return head
	}
	// 获取链表长度
	current := head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	// 边界条件：n > 链表长度
	if n > length {
		return head
	}
	// 边界条件：n 等于链表长度 删除头节点
	if n == length {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}

	current = head // 这里是将链表的地址赋值给了
	i := 0
	for current != nil {
		if i == length-n-1 {
			// 链表移动
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	// 链表是引用类型 对 current 的操作都会体现在 head 上
	return head
}
