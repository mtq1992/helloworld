package main

import (
	"fmt"
)

// Circle /* 定义结构体 */
type Circle struct {
	radius float64
	area   float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
	c1.getTestArea()
	fmt.Println(c1.area)
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

// 该 method 属于 Circle对象指针中的方法
func (c *Circle) getTestArea() {
	c.area = 3.14 * c.radius * c.radius
}
