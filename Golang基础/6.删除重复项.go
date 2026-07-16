package main

// 删除重复项
func removeDuplicates(nums []int) int {
	// 创建map 记录每个元素出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	// 清空切片
	nums = nums[:0]
	// 遍历key 到切片
	for k := range m {
		nums = append(nums, k)
	}
	return len(nums)
}
