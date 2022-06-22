package structures

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int       // 数据域
	Next *ListNode // 指针域
}

// List2Ints 链表转换为Int数组
func List2Ints(head *ListNode) []int {
	// 链条深度限制
	limit := 100
	times := 0

	var res []int
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List Int数组转链表
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{} // 初始化一个空链表节点
	t := l           // 取头节点
	for _, v := range nums {
		t.Next = &ListNode{Val: v} // 头结点的指针域指针(头指针)指向下一节点
		t = t.Next                 // 取下一节点 继续进行链表结构构建(指针赋值)
	}
	// 返回最终构建的链表
	return l.Next
}

// GetNodeWith 链表操作-获取节点
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Ints2ListWithCycle 循环链表
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	// pos = -1 时 不构建循环
	if pos == -1 {
		return head
	}
	c := head

	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
