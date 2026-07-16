package main

// 给定一个只包含 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(str string) bool {
	// 创建切片
	var stack []string
	for i := 0; i < len(str); i++ {
		// 字符顺序全部添加到切片
		stack = append(stack, string(str[i]))
	}
	// 创建判断数组
	var valid = []string{"()", "{}", "[]"}
	for i := 0; i < len(stack)-1; i++ {
		// 和下一个匹配是否一致
		var isTrueNext bool
		// 和上一个匹配是否一致
		var isTruePre bool
		// 判断字符串拼接是否匹配
		var bracketPre string
		bracketNext := stack[i] + stack[i+1]
		if i > 0 {
			bracketPre = stack[i-1] + stack[i]
		}
		for s := 0; s < len(valid); s++ {
			if bracketNext == valid[s] {
				// 如果匹配上了，isTrue为true
				isTrueNext = true
				break
			}
			if bracketPre == valid[s] {
				// 如果匹配上了，isTrue为true
				isTruePre = true
				break
			}
		}
		// 判断如果匹配上就删除
		if isTrueNext {
			stack = append(stack[:i], stack[i+2:]...)
			i = -1
			continue // 跳出当前循环，重新判断
		}
		// 判断如果匹配上就删除
		if isTruePre {
			stack = append(stack[:i-1], stack[i+1:]...)
			i = -1
			continue // 跳出当前循环，重新判断
		}
	}

	// 判断最终切片长度是否时0，如果是0就是所有的都匹配上了
	return len(stack) == 0
}
