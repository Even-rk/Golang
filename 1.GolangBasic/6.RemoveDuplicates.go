package main

// 删除重复项，保留原有的顺序,入参使用指针，需要操作原切片
func removeDuplicates(nums *[]int) int {
	// 遍历切片
	for i := 0; i < len(*nums)-1; i++ {
		// 判断当前元素是否是重复的
		if (*nums)[i] == (*nums)[i+1] {
			// 删除重复的元素
			*nums = append((*nums)[:i], (*nums)[i+1:]...)
			i = -1 // 跳出当前循环，重新判断
			continue
		}
	}

	return len(*nums)
}
