package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 先建立一个虚拟头结点 Next 指向真正的 Head
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			// 获取 l1 链表当前节点的值
			n1 = l1.Val
			// 移动节点到下一节点
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			// 获取 l2 链表当前节点的值
			n2 = l2.Val
			// 移动节点到下一个节点
			l2 = l2.Next
		}
		// 相当于是从个位向前取和进位
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
