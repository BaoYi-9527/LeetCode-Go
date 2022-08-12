package leetcode

// intToRoman
// @Description: 暴力枚举
// @param num
// @return string
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
				currentStr = intMap[i] + currentStr
				remainder--
			}
		}
		romanStr = currentStr + romanStr
	}
	return romanStr
}

func intToRoman1(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		// 寻找到 num 的最大位
		for values[i] > num {
			i++
		}
		// 依次将其最大位对应的罗马数字放在左边 并对数值进行减小继续遍历
		num -= values[i]
		res += symbols[i]
	}
	return res
}
