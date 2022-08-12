package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	resMap := map[string]int{
		"III":     3,
		"IV":      4,
		"LVIII":   58,
		"MCMXCIV": 1994,
	}
	for x, res := range resMap {
		result := romanToInt(x)
		if res != result {
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
