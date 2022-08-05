package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	resMap := map[string]string{
		"babad":      "aba", // bab 也是正确答案
		"cbba":       "bb",
		"babcddcbae": "abcddcba",
		"aaaa":       "aaaa",
		"b":          "b",
	}

	for str, res := range resMap {
		//result := longestPalindrome(str)
		result := longestPalindrome1(str)
		if res != result {
			fmt.Println("string: ", str)
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
