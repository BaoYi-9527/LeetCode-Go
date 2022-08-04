package leetcode

import (
	"strings"
)

// lengthOfLongestSubstring
// @Description: 无无重复字符的最长子串【暴力破解法】
// @param s
// @return int
func lengthOfLongestSubstring(s string) int {
	strArr := strings.Split(s, "") // 现将字符串拆分为数组
	maxLen := 0
	curStrArr := make([]string, 0) // 创建一个数组用于存储无重复子串
	for _, value := range strArr { // 循环字符串
		// 边界条件 第一次进入
		if len(curStrArr) == 0 {
			curStrArr = append(curStrArr, value) // 第一个字符
		} else {
			for key, item := range curStrArr { // 循环子串
				// 若子串中有字符与当前字符重复则截取子串重复字符后面的子串
				if item == value {
					curStrArr = curStrArr[(key + 1):]
				}
			}
			// 将当前字符加入到无重复子串的后面
			curStrArr = append(curStrArr, value)
		}
		// 不断更新 最长子串长度 若当前子串长度大于之前的就更新
		if len(curStrArr) > maxLen {
			maxLen = len(curStrArr)
		}
	}
	return maxLen
}

// 解题思路：
// a b c => 3
// a b c a => b c a => 3
// b c a b => c a b => 3
// c a b c => a b c => 3
// a b c b => c b => 2
// c b b => b => 1

// lengthOfLongestSubstring1
// @Description: 位图
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 这里解释一下为什么 bitSet 作为一个数组类型可以存储 s[right]
		// 比如 s = "abc"  那么 s[0] 是多少呢，用PHP的话我们可能下意识就觉得会是 “a”
		// 但是在 Golang 中，对字符串取下标只对纯ASCII码的字符串有效且取出的是字符的ASCII码
		// 也就是 s[0] = 97 所以取出的是一个 int 类型的值，完全可以作为一个数组的下标

		// 若 bitSet 中存在 s[right] 字符，则 bitSet 最左边的字符设置为 false
		// 思路在于 不管当前的字符是和无重复子串的第一个重复的还是中间部分的字重复的，首部的字符都将被舍弃
		// 例：无重复字符串 abcd 若循环中的元素为 d 则 d 前面的所有子串的都是必然需要被舍弃的
		// d第一次出现后 right没有变 下次循环依然为 d 也就是一直会循环到 left 也为 d  新的无重复子串为 d
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+right >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}
