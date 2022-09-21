package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	resMap := map[string][]string{
		"fl": []string{"flower", "flow", "flight"},
		"":   []string{"dog", "racecar", "car"},
	}

	for res, strs := range resMap {
		result := longestCommonPrefix(strs)
		if res != result {
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
