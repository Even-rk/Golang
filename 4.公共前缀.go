package main

func longestCommonPrefix(strs []string) string {
	// 如果数组元素在1个以下，直接返回空字符串
	if len(strs) <= 0 {
		return ""
	}
	// 最小的字符串长度
	minLen := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}
	//初始化切片保存公共前缀，默认为最小长度字符串组成的切片
	var prefix []string
	// 循环第一个字符串
	for s := 0; s < minLen; s++ {
		prefix = append(prefix, string(strs[0][s]))
	}
	// 循环需要验证的字符串组
	for i := 0; i < len(strs); i++ {
		// 以最小长度为基准，开始循环
		for s := 0; s < minLen; s++ {
			// 判断第i个字符串的第s个字符 和 切片中的第s个元素 不一致
			if string(strs[i][s]) != prefix[s] {
				// 不一致的设置为 空字符串
				prefix[s] = ""
			}
		}
	}
	// 最终结果切片组合为字符串
	var commonPrefix string
	for i := 0; i < len(prefix); i++ {
		commonPrefix += prefix[i]
	}
	return commonPrefix
}
