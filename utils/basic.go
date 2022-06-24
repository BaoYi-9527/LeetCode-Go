package utils

import "fmt"

// Recursion 阶乘
func Recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursion(n-1)
}

// RecursionTail 尾递归
func RecursionTail(n int, a int) int {
	if n == 1 {
		return a
	}
	// 尾递归在调用自身后直接返回了其值，在后续递归中不会对之前的步骤进行再计算
	return RecursionTail(n-1, a*n)
}

// Fibonacci 斐波那契数列
func Fibonacci(n, a1, a2 int) int {
	fmt.Println(a2)
	if n == 0 {
		return a2
	}

	return Fibonacci(n-1, a2, a1+a2)
}

// BinarySearch 二分查找
func BinarySearch(array []int, target int, left, right int) int {
	// 边界条件 找不到
	if left > right {
		return -1
	}
	// 二分二分 从中间开始分
	mid := (left + right) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid // 正好找到了
	} else if middleNum > target {
		return BinarySearch(array, target, 0, mid-1)
	} else {
		return BinarySearch(array, target, mid+1, right)
	}
}
