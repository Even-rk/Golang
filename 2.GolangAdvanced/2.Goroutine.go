package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 函数 打印从1到10的奇数
func printOdd() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

// 函数 打印从2到10的偶数
func printEven() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// 执行调用，协程启动
func printOddAndEven() {
	go printOdd()
	go printEven()
}

// 2.设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 任务执行函数
func taskExecutor(task func(), num int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 第num个任务开始时间
	fmt.Printf("第%d个任务开始时间: %v\n", num, time.Now())
	// 执行任务
	task()
	// 第num个任务结束时间
	fmt.Printf("第%d个任务结束时间: %v\n", num, time.Now())
}

// 任务调度器
func taskScheduler(tasks []func()) {
	// 初始化任务列
	var wg sync.WaitGroup
	for i := 0; i < len(tasks); i++ {
		// 任务列➕1
		wg.Add(1)
		go taskExecutor(tasks[i], i, &wg)
	}
	// 等待所有任务完成
	wg.Wait()
}
