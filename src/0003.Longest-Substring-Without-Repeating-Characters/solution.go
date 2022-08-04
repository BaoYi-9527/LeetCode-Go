package leetcode

import "strings"

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
