package leetcode

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	resMap := map[int]bool{
		121:    true,
		-121:   false,
		100001: true,
		10:     false,
		0:      true,
	}
	for x, res := range resMap {
		result := isPalindrome(x)
		if res != result {
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
