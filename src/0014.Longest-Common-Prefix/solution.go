package leetcode

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
