package leetcode

// @Description:最长回文子串

// longestPalindrome
// @Description:动态规划解决
// @param s
// @return string
func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	// a b c d d c b a
	// 0 1 2 3 4 5 6 7
	// i = 5, j =7
	// i 从末尾向前遍历
	for i := len(s) - 1; i >= 0; i-- {
		// j 从 i 的位置向后遍历
		for j := i; j < len(s); j++ {
			// dp[i][j] 的值标识 从 i -> j 的子串是否为一个回文子串
			// (s[i] == s[j]) 表示头尾相同
			// (j - i < 3) 是为了兼容 aba/aa/a 这种情况
			// dp[i + 1][j - 1] 是表示当前的回文子串去除头尾后依然是一个回文子串
			// 值得注意的是 在 i 处于最后 3 个位置时，||条件短路 不会走到 dp[i + 1][j - 1]
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			// 如果当前子串是一个回文子串 且当前子串的长度大于存储的最长回文子串的长度就更新最长回文子串
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// longestPalindrome1
// @Description: 窗口解法
// @param s
// @return string
func longestPalindrome1(s string) string {
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
