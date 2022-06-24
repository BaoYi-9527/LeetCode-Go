package main

import (
	"LeetcodeGo/utils"
	"fmt"
)

func main() {
	basic()
}

func basic() {
	//fmt.Println(utils.Recursion(10))	// 递归
	//fmt.Println(utils.RecursionTail(10, 1))	// 尾部递归
	//utils.Fibonacci(10, 1, 1)	// 斐波那契数列
	arr := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	// 下标为数组长度-1
	result := utils.BinarySearch(arr, 123, 0, len(arr)-1) // 二分查找
	fmt.Println(result)
}
