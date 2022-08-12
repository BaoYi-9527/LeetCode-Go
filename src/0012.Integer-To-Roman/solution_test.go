package leetcode

import "testing"

func TestIntToRoman(t *testing.T) {
	resMap := map[int]string{
		3:    "III",
		4:    "IV",
		9:    "IX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}
	for x, res := range resMap {
		result := intToRoman(x)
		if res != result {
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
