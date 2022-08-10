package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		items := make([]int, 0)
		for {
			item := x % 10
			items = append(items, item)
			x = x / 10
			if x == 0 {
				break
			}
		}
		count := len(items)
		isOdd := count % 2
		middle := count/2 - 1
		i := 1
		for middle >= 0 {
			// 偶数
			if isOdd == 0 && items[middle] != items[middle+i] {
				return false
			}
			// 奇数
			if isOdd != 0 && items[middle] != items[middle+i+1] {
				return false
			}
			i = i + 2
			middle--
		}
		return true
	}
}

func isPalindrome1(x int) bool {
	// 边界条件
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	// 将数分解为数组
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	// 头尾对比
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}
