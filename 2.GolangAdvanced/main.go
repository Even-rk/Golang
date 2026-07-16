package main

import "fmt"

func main() {
	// 一.指针
	// 1.
	num := 20
	addTen(&num)
	fmt.Println(num)
	// 2.
	slice := []int{1, 2, 3}
	doubleSlice(&slice)
	fmt.Println(slice)
	// 二.Goroutine
}
