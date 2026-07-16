package main

func isPalindrome(x int) bool {
	// 如果小于等于0，直接返回false
	if x < 0 {
		return false
	}
	// 如果是0，直接返回true
	if x == 0 {
		return true
	}
	// 记录x
	original := x
	// 倒转数字
	var reverse int
	for x > 0 {
		reverse = reverse*10 + x%10 // 每次把最后一位拿出来放到reverse后面
		x /= 10                     // 去掉最后一位
	}

	// 判断是否相等
	return reverse == original
}
