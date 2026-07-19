package main

import "sort"

// 合并区间
func merge(intervals [][]int) [][]int {
	// 导入sort包后使用标准库排序，按区间的起始位置升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 遍历切片
	for i := 0; i < len(intervals)-1; i++ {
		// 判断前一个区间的结束位和后一个区间的开始位
		// 创建新的区间，开始位使用前一个区间的开始，结束位使用后一个位区间的结束
		// 1，前者大于后者， 2，前者等于后者

		// 前者的开始位和结束位
		preStart := intervals[i][0]
		preEnd := intervals[i][1]
		// 后者的开始位和结束位
		nextStart := intervals[i+1][0]
		nextEnd := intervals[i+1][1]
		// 判断 是否符合 12条件
		if preEnd > nextStart || preEnd == nextStart {
			// 符合条件，创建一个新的区间
			newInterval := []int{preStart, max(preEnd, nextEnd)}
			// 合并区间, 利用切片删除这两个区间在在原来的位置把新的区间添加
			intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i+2:]...)...)
			i = -1 // 跳出当前循环，重新判断
			continue
		}
	}
	// 返回合并后的区间
	return intervals
}
