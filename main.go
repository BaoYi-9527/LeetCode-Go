package main

import (
	"fmt"
)

func main() {
	//s := "PAYPALISHIRING"
	// 最长回文子串
	//fmt.Println(longestPalindrome(s))
	// Z 字形变换
	//fmt.Println(convert(s, 4))
	// 整数反转
	x := -123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	tmp := 0
	for x != 0 {
		// tmp 从 x 的个位数开始每次进一位(扩大10)
		tmp = tmp*10 + x%10
		// x 从个位数开始每次退一位(缩小10)
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func convert(s string, numRows int) string {
	// matrix 二维数组用于z形结构
	// down 表示了Z第一结构上的元素位置[行]，也就是当前元素属于第几行
	// up 为拐角元素所属的列 numRows - 2 (2 表示的第一行和最后一行)
	// 也就是 up 是从倒数第二行开始向上走
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		// P	 	 	I			N
		// A	 	L	S		I	G
		// Y	A		H	R
		// P			I
		// down != numRows 时 循环的是 Z 字的第一结构
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 { // down = numRows 后 进入拐角后的元素
			// 此时up是由下向上走(行)，直到 up 走到 0 走完了，也就是拐角结束了
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			// 重置循环 但元素遍历不进入下一个元素
			// 只是重置了 up 和 down
			// 需要注意的是 循环并不是以一个完整的 Z 形结构为一个循环的
			// 而是以 不包含底行的 Z 行结构(7形结构)为一个循环
			up = numRows - 2
			down = 0
		}
	}
	// solution 存储结果数组
	solution := make([]byte, 0, len(s))
	// 将 matrix 矩形结构的二维数组按行顺序扁平为一维数组即为需要得到的元素的数组排序
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	// 数组转字符串
	return string(solution)
}

func longestPalindrome(s string) string {
	// a b c d d c b a
	// 0 1 2 3 4 5 6 7

	// left, right 存储中心两侧的位置下标
	// pl, pr 存储最长回文子串的起始和结束位置
	//left, right, pl, pr := 0, -1, 0, 0
	left, right, pl, pr := 0, 0, 0, 0
	for left < len(s) {
		// 当 left = right+1 时 回文子串为一个偶数串
		// right 初始值为 -1 的意义就在于
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// left 为中心节点左边的位置 right 为中心节点右边的位置
		// 中心节点不一定存在 也就是当回文子串为一个偶数子串是 中心节点为一个虚拟节点
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		// 当回文子串的长度 大于存储的最长回文子串的长度时更新最长回文子串的下标
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// left 和 right 每次重置为 中心节点右边的节点
		// 也就是 left 和 right 的位置起始是中心节点的位置
		left = (left+right)/2 + 1
		right = left
	}
	// 需要注意的是切片操作时是不包含最右边的元素的 所以需要 pr+1
	return s[pl : pr+1]
}
