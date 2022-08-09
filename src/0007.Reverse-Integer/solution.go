package leetcode

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
