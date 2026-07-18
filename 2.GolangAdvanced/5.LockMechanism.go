package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 1.编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func runMutex() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			// 协程递增
			defer wg.Done()
			fmt.Println(i)
			for j := 0; j < 1000; j++ {
				// 1. 加锁
				mu.Lock()
				counter++
				mu.Unlock() //  在每次循环结束时立即解锁
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("最终计数器值:", counter)
}

// 2.使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func runAtomic() {
	// 使用int64类型，原子操作需要
	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			// 协程递增
			defer wg.Done()
			fmt.Println(i)
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("最终计数器值:", counter)
}
