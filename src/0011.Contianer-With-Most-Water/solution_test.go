package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := maxArea(input)
	correct := 49
	if output != correct {
		t.Errorf("【expected】:%v | 【actual】:%v\n", correct, output)
	}

	input = []int{1, 1}
	output = maxArea(input)
	correct = 1
	if output != correct {
		t.Errorf("【expected】:%v | 【actual】:%v\n", correct, output)
	}

}
