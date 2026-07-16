package main

// 非空数组，除了某些元素只出现一次之外，其余每个元素均出现两次。找出那个只出现一次的元素。
func findSingle(nums []int) []int {
	// 创建map
	var mapNums = make(map[int]int)
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 统计出现次数写入到map中
		mapNums[nums[i]] = mapNums[nums[i]] + 1
	}
	// 创建切片
	var single []int
	// 遍历map中的元素，取出现次数为1的元素
	for k, v := range mapNums {
		if v == 1 {
			single = append(single, k)
		}
	}
	// 返回切片结果【返回所有次数为1的元素】
	return single
}
