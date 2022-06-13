package leetcode

func twoSum(nums []int, target int) []int {
	// 创建一个 map 用于存储 nums 中的 [值 => 键] 对应关系
	m := make(map[int]int)
	// 遍历 nums 中的键值对
	for k, v := range nums {
		// 如果目标值 target - 当前值 的结果存在于 map 中
		// 则是我们想要的结果 可以直接在 map 中通过值取出对应的下标
		// 而当前值的下标就在此次遍历中可以得到
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		// 将 nums 的键值关系反转后存储在 map 中
		m[v] = k
	}
	// 不存在
	return nil

	// 解题思路
	// 在遍历原有数组的过程中不断的将数组的键值关系反转后存储到 map 中
	// 从而方便后续通过值直接得到其在数组中的下标
	// 当俩数之和的其中一个值进入 map 中后，遍历第二个值时就可以在 map 中找到第一个值
	// 从而返回俩个数的下标
}
