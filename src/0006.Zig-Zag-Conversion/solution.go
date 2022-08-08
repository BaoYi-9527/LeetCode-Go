package _006_Zig_Zag_Conversion

func convert(s string, numRows int) string {
	// matrix 二维数组用于z形结构
	// down 表示了Z第一结构上的元素位置[行]，也就是当前元素属于第几行
	// up 为拐角元素所属的列 numRows - 2 (2 表示的第一行和最后一行)
	// 也就是 up 是从倒数第二行开始向上走
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		// 0 P	 	 	I			N
		// 1 A	 	L	S		I	G
		// 2 Y	A		H	R
		// 3 P			I
		// down != numRows 时 循环的是 Z 字的第一结构
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 { // down = numRows 后 进入拐角后的元素
			// 此时up是由下向上走(行)，直到 up 走到 0 走完了，也就是拐角结束了
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			// 重置循环 但元素遍历不进入下一个元素
			// 只是重置了 up 和 down
			// 需要注意的是 循环并不是以一个完整的 Z 形结构为一个循环的
			// 而是以 不包含底行的 Z 行结构(7形结构)为一个循环
			up = numRows - 2
			down = 0
		}
	}
	// solution 存储结果数组
	solution := make([]byte, 0, len(s))
	// 将 matrix 矩形结构的二维数组按行顺序扁平为一维数组即为需要得到的元素的数组排序
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	// 数组转字符串
	return string(solution)
}
