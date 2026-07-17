package main

import (
	"fmt"
	"math"
)

// 1.定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// Shape接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Area 实现Shape接口的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 实现Shape接口的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area 实现Shape接口的Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 实现Shape接口的Perimeter方法 - 圆周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 2.使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 创建Person实例
type Person struct {
	Name string
	Age  int
}

// 继承Person实例，并新添加EmployeeID字段
type Employee struct {
	Person
	EmployeeID int
}

// 打印员工信息
func (e Employee) PrintInfo() {
	fmt.Printf(e.Name, e.Age, e.EmployeeID)
}
