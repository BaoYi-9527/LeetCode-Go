package main

import "fmt"

func main() {
	//s := "PAYPALISHIRING"
	// 最长回文子串
	//fmt.Println(longestPalindrome(s))
	// Z 字形变换
	//fmt.Println(convert(s, 4))
	// 整数反转
	//x := -123
	//fmt.Println(reverse(x))
	//x := 1001
	//fmt.Println(isPalindrome(x))
	//height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//fmt.Println(maxArea(height))
	//num := 2300
	//fmt.Println(intToRoman(num))
	//str := "MCMXCIV"
	//fmt.Println(romanToInt(str))
	strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	// 取第一个字符串为假定的最长公共前缀
	// 若该字符串恰好为最短的字符串且全为公共前缀则刚好
	// 若该字符串恰好为最长字符串或字符串仅前面一部分为公共前缀 则后续会对其进行截取
	prefix := strs[0]
	// 遍历字符串数组
	for i := 1; i < len(strs); i++ {
		// 遍历公共前缀
		for j := 0; j < len(prefix); j++ {
			// 当前字符串的长度小于公共前缀遍历的长度时 或当前字符串遍历的字符与公共前缀的字符不一致时
			// 说明当前 j 即为最长字符的下标
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}

func romanToInt(s string) int {
	symbolsMap := map[string]int{
		"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
	}
	max, res, i := 0, 0, len(s)-1
	for i >= 0 {
		current := symbolsMap[string(s[i])]
		if current >= max {
			res = res + current
			max = current
		} else {
			res = res - current
		}
		i--
	}
	return res
}

func intToRoman(num int) string {
	// 1 < num < 3999
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	intMap := map[int]string{
		1: "I",
		2: "X",
		3: "C",
		4: "M",
	}
	intArr := []int{1, 10, 100}
	romanStr := ""

	for i := 1; i <= 4; i++ {
		currentStr := ""
		if i < 4 {
			remainder := num % 10
			num = num / 10
			if remainder > 1 && remainder < 4 {
				for remainder > 0 {
					currentStr = intMap[i] + currentStr
					remainder--
				}
			} else if remainder > 5 && remainder < 9 {
				currentStr = romanMap[5*intArr[i-1]] + currentStr
				remainder = remainder - 5
				for remainder > 0 {
					currentStr = currentStr + intMap[i]
					remainder--
				}
			} else {
				currentStr = romanMap[remainder*intArr[i-1]] + currentStr
			}
		} else {
			remainder := num % 10
			num = num / 10
			for remainder > 0 {
				currentStr = "M" + currentStr
				remainder--
			}
		}
		romanStr = currentStr + romanStr
	}
	return romanStr
}

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		// 左边小就左边移动 右边小就右边移动
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		// 计算面积 更新最大面积
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

func isPalindrome(x int) bool {
	// 边界条件
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	// 将数分解为数组
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	// 头尾对比
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
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
