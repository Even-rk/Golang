package main

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
func plusOne(digits []int) []int {
	// 倒序遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果是9就设置为0、
		if digits[i] == 9 && i != 0 {
			digits[i] = 0
		} else if i == 0 && digits[i] == 9 {
			// 如果是最高位，且是9，改为0并新增一个位，值为1
			digits[i] = 0
			digits = append([]int{1}, digits...)
		} else {
			// 如果不是9，就加1
			digits[i]++
			// 加1的动作只执行一次就跳出
			break
		}
	}
	return digits
}
