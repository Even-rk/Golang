package main

import "fmt"

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	fmt.Println(nums, target)
	// 初始化结果切片
	var result []int
	// 遍历切片
	for i := 0; i < len(nums); i++ {
		// 遍历切片的下一个元素
		for j := i + 1; j < len(nums); j++ {
			// 如果两个元素的和等于目标值
			if nums[i]+nums[j] == target {
				// 将索引下标添加到结果切片
				result = append(result, i, j)
				break
			}
		}
	}

	return result
}
