package main

import (
	"fmt"
	"sync"
)

// 1.编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func sendNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// 发送消息到通道
		ch <- i
	}
}

// receiveNumbers 接收通道中的消息并打印出来
func receiveNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// 接收通道中的消息
		fmt.Println(<-ch)
	}
}

// 执行函数
func runChannel() {
	// 创建一个WaitGroup，用于等待所有子协程完成
	var wg sync.WaitGroup
	wg.Add(2)
	// 创建一个channel
	ch := make(chan int)
	// 启动sendNumbers协程发送消息
	go sendNumbers(ch, &wg)
	// 启动receiveNumbers协程接受消息
	go receiveNumbers(ch, &wg)
	// 等待所有子协程完成
	wg.Wait()
}

// 2.实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func sendNumbersWithBuffer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		// 发送消息到通道
		ch <- i
	}
}

func receiveNumbersWithBuffer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		// 接收通道中的消息
		fmt.Println(<-ch)
	}
}

func runChannelWithBuffer() {
	// 创建一个channel，缓冲区为5
	ch := make(chan int, 5)
	// 创建一个WaitGroup，用于等待所有子协程完成
	var wg sync.WaitGroup
	wg.Add(2)
	// 启动sendNumbersWithBuffer协程发送消息
	go sendNumbersWithBuffer(ch, &wg)
	// 启动receiveNumbersWithBuffer协程接受消息
	go receiveNumbersWithBuffer(ch, &wg)
	// 等待所有子协程完成
	wg.Wait()
}
