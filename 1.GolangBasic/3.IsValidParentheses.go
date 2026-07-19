package main

// 给定一个只包含 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(str string) bool {
	// 初始化切片用于存储未匹配的左括号
	var leftVals []string
	// 建立右括号到对应左括号的映射，用于快速匹配
	rightKey := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	// 遍历字符
	for _, val := range str {
		// 当前字符
		crts := string(val)
		// 如果是右括号，看切片中是否有匹配的左括号
		if left, ok := rightKey[crts]; ok {
			// 切片不为空且切片的最后一个元素是匹配的左括号，删除切片最后一个元素
			if len(leftVals) > 0 && leftVals[len(leftVals)-1] == left {
				// 删除切片最后一个元素
				leftVals = leftVals[:len(leftVals)-1]
			} else {
				// 不匹配，直接返回无效
				return false
			}
		} else {
			// 如果是左括号，直接放到切片中
			leftVals = append(leftVals, crts)
		}
	}
	// 遍历结束后看切片是否为空，为空则所有括号都匹配成功
	return len(leftVals) == 0
}
