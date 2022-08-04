package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	strToResMap := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}
	for str, res := range strToResMap {
		//result := lengthOfLongestSubstring(str)
		result := lengthOfLongestSubstring1(str)
		if res != result {
			fmt.Println("string: ", str)
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
