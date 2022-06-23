package structures

import "fmt"

// DLinkedList 双向链表
type DLinkedList struct {
	head, tail *DLinkedNode
	size       int // 长度和数据都不包含头尾节点
}

// DLinkedNode 双向链表节点
type DLinkedNode struct {
	pre, next *DLinkedNode
	key, val  interface{}
}

// NewDLinkList 初始化一个链表
func NewDLinkList() *DLinkedList {
	list := &DLinkedList{
		head: new(DLinkedNode),
		tail: new(DLinkedNode),
	}
	// 形成循环链表
	list.head.next = list.tail
	list.tail.pre = list.head
	return list
}

// NewDLinkNode 初始化一个链表节点
func NewDLinkNode(key interface{}, val interface{}) *DLinkedNode {
	return &DLinkedNode{key: key, val: val}
}

// Empty 链表方法-链表是否为空
func (list *DLinkedList) Empty() bool {
	if list.head.next.next == nil && list.head.next.key == nil && list.head.next.val == nil {
		return true
	} else {
		return false
	}
}

// RemoveNode 链表方法-移除一个节点
func (list *DLinkedList) RemoveNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	list.size--
}

// AddToHead 链表方法-头部新增一个节点
func (list *DLinkedList) AddToHead(node *DLinkedNode) {
	// 指定新增节点的前后节点
	node.pre = list.head
	node.next = list.head.next
	// 修改前后节点的指针
	list.head.next.pre = node
	list.head.next = node
	list.size++
}

// AddToTail 链表方法-尾部部新增一个节点
func (list *DLinkedList) AddToTail(node *DLinkedNode) {
	// 指定新增节点的前后节点
	node.next = list.tail
	node.pre = list.tail.pre
	// 修改前后节点的指针
	list.tail.pre.next = node
	list.tail.pre = node
	list.size++
}

// RemoveHead 链表方法-头部移除节点
func (list *DLinkedList) RemoveHead(node *DLinkedNode) *DLinkedNode {
	next := list.head.next
	list.RemoveNode(next)
	return next // 返回删除的节点
}

// RemoveTail 链表方法-尾部移除节点
func (list *DLinkedList) RemoveTail(node *DLinkedNode) *DLinkedNode {
	pre := list.tail.pre
	list.RemoveNode(pre)
	return pre
}

// MoveToHead 链表方法-节点移动至头部
func (list *DLinkedList) MoveToHead(node *DLinkedNode) {
	list.RemoveNode(node)
	list.AddToHead(node)
}

// MoveToTail 链表方法-节点移动至尾部
func (list *DLinkedList) MoveToTail(node *DLinkedNode) {
	list.RemoveNode(node)
	list.AddToTail(node)
}

// Length 返回链表的长度
func (list *DLinkedList) Length() int {
	return list.size
}

func (list *DLinkedList) Print() {
	cur := list.head.next
	for cur != list.tail {
		fmt.Print("{key:", cur.key, "val:", cur.val, "} => ")
		cur = cur.next // 指向下一个节点
	}
	fmt.Println()
}
