package leetcode

import (
	"LeetcodeGo/utils"
	"fmt"
	"testing"
)

// IntArr2List Int数组转链表
func IntArr2List(nums []int) *ListNode {
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

// List2IntArr 链表转换为Int数组
func List2IntArr(head *ListNode) []int {
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

func TestAddTwoNumbers(t *testing.T) {
	// 测试用例①
	l1 := IntArr2List([]int{2, 4, 3})
	l2 := IntArr2List([]int{5, 6, 4})
	expect := []int{7, 0, 8}
	calcRes := addTwoNumbers(l1, l2)
	calcResArr := List2IntArr(calcRes)

	if utils.IntArrayEquals(expect, calcResArr) {
		fmt.Printf("【output】:%v\n", calcResArr)
	} else {
		fmt.Printf("【expected】:%v\n", expect)
		t.Errorf("【expected】:%v | 【actual】:%v\n", expect, calcResArr)
	}

	// 测试用例②
	l1 = IntArr2List([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = IntArr2List([]int{9, 9, 9, 9})
	expect = []int{8, 9, 9, 9, 0, 0, 0, 1}
	calcRes = addTwoNumbers(l1, l2)
	calcResArr = List2IntArr(calcRes)

	if utils.IntArrayEquals(expect, calcResArr) {
		fmt.Printf("【output】:%v\n", calcResArr)
	} else {
		fmt.Printf("【expected】:%v\n", expect)
		t.Errorf("【expected】:%v | 【actual】:%v\n", expect, calcResArr)
	}

}
