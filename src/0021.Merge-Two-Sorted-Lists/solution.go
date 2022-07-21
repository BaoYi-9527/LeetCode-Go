package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 若是 list1 当前节点值 < list2 当前节点值
	// 则说明 list2 当前节点值 排序至少在 list1 当前节点值都后面
	// 所以递归继续拿 list1.Next 的值与 list2 当前节点的值进行比较
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
