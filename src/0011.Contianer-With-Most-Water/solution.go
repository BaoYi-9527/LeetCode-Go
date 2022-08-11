package leetcode

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		// 左边小就左边移动 右边小就右边移动
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		// 计算面积 更新最大面积
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
