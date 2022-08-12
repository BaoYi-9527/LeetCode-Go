package leetcode

// romanToInt
// @Description: 简单遍历
// @param s
// @return int
func romanToInt(s string) int {
	symbolsMap := map[string]int{
		"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1,
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
