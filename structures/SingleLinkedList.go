package structures

import "fmt"

// SingleNode 单链表节点
type SingleNode struct {
	data interface{}
	next *SingleNode
}

// SingleLinkedList 单链表
type SingleLinkedList struct {
	head *SingleNode
}

func NewSingleLinkedList(data interface{}) *SingleLinkedList {
	return &SingleLinkedList{NewSingleNode(data)}
}

func NewSingleNode(data interface{}) *SingleNode {
	return &SingleNode{data: data}
}

// Empty 链表是否为空
func (list *SingleLinkedList) Empty() bool {
	if list.head == nil {
		return true
	} else {
		return false
	}
}

// Length 链表长度
func (list *SingleLinkedList) Length() int {
	cur := list.head // 取头节点游标
	count := 0
	for {
		if cur == nil {
			break
		} else {
			count++
			cur = cur.next
		}
	}
	return count
}

// Unshift 从开头插入一个节点
func (list *SingleLinkedList) Unshift(data interface{}) {
	// 链表头部插入
	node := &SingleNode{data: data}
	// 将新的节点插入到链表头节点的前面 成为新的头节点
	// temp := list.head
	// node.next = temp	 新节点的下一个节点是之前的头节点
	// list.head = node	 链表新的头节点是插入的节点
	node.next, list.head = list.head, node
}

// Push 从尾部插入一个节点
func (list *SingleLinkedList) Push(data interface{}) {
	// 链表尾部插入
	node := &SingleNode{data: data}

	if list.Empty() {
		list.head = node
	} else {
		cur := list.head
		for {
			if cur.next != nil {
				cur = cur.next
			} else { // cur.next == nil 尾部节点 尾部节点的指针指向新增节点
				cur.next = node
				break
			}
		}
	}
}

// Insert 指定位置插入一个节点
func (list *SingleLinkedList) Insert(data interface{}, pos int) {
	if pos <= 0 {
		list.Unshift(data)
	} else if pos >= list.Length() {
		list.Push(data)
	} else {
		node := &SingleNode{data: data}
		count := 0 // 标记位置
		pre := list.head
		for count < (pos - 1) {
			count++
			pre = pre.next
		}
		// 退出循环表示找到位置
		node.next = pre.next
		pre.next = node
	}
}

// Remove 删除指定值的节点
func (list *SingleLinkedList) Remove(data interface{}) {
	cur := list.head
	if cur.data == data {
		// 删除头结点
		list.head = cur.next
	}
	for cur.next != nil {
		if cur.next.data != data {
			cur = cur.next
		} else {
			cur.next = cur.next.next
		}
	}
}

// LocateElem 查找值为 data 的数据 返回其下标
func (list *SingleLinkedList) LocateElem(data interface{}) int {
	cur := list.head
	count := 0
	index := -1 // 不存在返回 -1
	for {
		if cur == nil { // 链表到底了
			fmt.Println("list don't have the data...")
			break
		} else if cur.data == data { // 找到了
			index = count
		} else { // 下一个节点
			cur = cur.next
		}
		count++
	}
	return index
}

// GetElem 获取指定位置的节点的值
func (list *SingleLinkedList) GetElem(pos int) interface{} {
	if pos <= 0 {
		return list.head.data
	} else if pos >= list.Length() {
		fmt.Println("Out of list length...")
		return nil
	} else {
		count := 0 // 标记位置
		cur := list.head
		for count < (pos - 1) {
			count++
			cur = cur.next
		}
		// 跳出循环时 cur 指向的即为指定位置得节点
		return cur.data
	}
}
