package main

// 合并区间
func merge(intervals [][]int) [][]int {
	// 先做一次排序
	// 谁的开始位最小，谁在前
	for i := 0; i < len(intervals)-1; i++ {
		// 如果当前区间的开始位置 大于 后一个区间的开始位, 两个区间交换
		if intervals[i][0] > intervals[i+1][0] {
			// 交换区间
			intervals[i], intervals[i+1] = intervals[i+1], intervals[i]
			// 交换后，回到第一个区间继续判断
			i = -1
			continue
		}
	}

	// 遍历切片
	for i := 0; i < len(intervals)-1; i++ {
		// 判断前一个区间的结束位和后一个区间的开始位
		// 创建新的区间，开始位使用前一个区间的开始，结束位使用后一个位区间的结束
		// 1，前者大于后者， 2，前者➕1 等于后者，3，前者等于后者

		// 前者的开始位和结束位
		preStart := intervals[i][0]
		preEnd := intervals[i][1]
		// 后者的开始位和结束位
		nextStart := intervals[i+1][0]
		nextEnd := intervals[i+1][1]
		// 判断 是否符合 123条件
		if preEnd > nextStart || preEnd+1 == nextStart || preEnd == nextStart {
			// 符合条件，创建一个新的区间
			newInterval := []int{preStart, nextEnd}
			// 合并区间, 利用切片删除这两个区间在在原来的位置把新的区间添加
			intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i+2:]...)...)
			i = -1 // 跳出当前循环，重新判断
			continue
		}
	}
	// 返回合并后的区间
	return intervals
}
