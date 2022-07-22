package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs
// @Description: 递归解决方法
// @param head
// @return *ListNode
func swapPairs(head *ListNode) *ListNode {
	// 边界结束条件 当链表中仅剩一个节点或没有节点时返回链表
	if head == nil || head.Next == nil {
		return head
	}
	h1, h2 := head, head.Next
	// 注意的是此处是 h2.Next 也就是自第三个节点开始的链表进行递归
	// 头节点指向递归后的链表 递归的子链表已在递归中完成了俩俩互换
	h1.Next = swapPairs(h2.Next)
	// 第二个节点指向投节点 完成首部的互换
	h2.Next = h1
	return h2
}

// swapPairs1
// @Description: 迭代解法
// @param head
// @return *ListNode
func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	l := make([]*ListNode, 0)

	// 将链表每隔 2 个切断存储到链表数组中
	for head != nil && head.Next != nil {
		l = append(l, head)
		head = head.Next.Next
	}
	// 结果链表
	nxt := (*ListNode)(nil)
	// 遍历链表数组长度
	// 注意这里对数组是逆向遍历的 从大到小 是为了保证最后组合的结果链表正确
	for i := len(l) - 1; i >= 0; i-- {
		// 俩俩交换
		h1, h2 := l[i], l[i].Next
		h1.Next = nxt // 第一个节点指向后续(已处理)链表 由于是逆向遍历 所以后面的链表依然在后面
		h2.Next = h1  // 第二个节点指向第一个节点
		nxt = h2      // 重新定义结果链表
	}

	return nxt
}
