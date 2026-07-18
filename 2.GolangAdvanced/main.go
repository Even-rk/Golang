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
	// 1.
	printOddAndEven()
	// 2.
	var tasks []func()
	tasks = append(tasks, func() { printOdd() })
	tasks = append(tasks, func() { printEven() })
	taskScheduler(tasks)
	// 三.面向对象
	// 1.
	// 创建实例
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	//调用方法
	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())
	fmt.Println(circle.Area())
	fmt.Println(circle.Perimeter())
	// 2.
	// 创建Employee实例
	employee := Employee{Person: Person{Name: "张三", Age: 30}, EmployeeID: 1001}
	// 调用实例方法
	employee.PrintInfo()
	// 四.channel
	// 1.
	runChannel()
}
