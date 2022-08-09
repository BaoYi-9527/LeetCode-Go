package leetcode

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	resMap := map[int]int{
		123:  321,
		-123: -321,
		120:  21,
		0:    0,
	}

	for x, res := range resMap {
		result := reverse(x)
		if res != result {
			fmt.Println("string: ", x)
			t.Errorf("【expected】:%v | 【actual】:%v\n", res, result)
		}
	}
}
