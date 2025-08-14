package main

import (
	"fmt"
	"math"
)

// Shape 接口定义
// 包含计算面积和周长的方法
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Rectangle 的字符串表示方法
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Circle 的字符串表示方法
func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// 通用函数：打印形状信息
func printShapeInfo(shape Shape) {
	fmt.Printf("形状: %v\n", shape)
	fmt.Printf("面积: %.2f\n", shape.Area())
	fmt.Printf("周长: %.2f\n", shape.Perimeter())
	fmt.Println("---")
}

// 通用函数：计算多个形状的总面积
func calculateTotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// 通用函数：计算多个形状的总周长
func calculateTotalPerimeter(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Perimeter()
	}
	return total
}

func main() {
	fmt.Println("=== Go面向对象编程示例 ===")
	fmt.Println("定义Shape接口，Rectangle和Circle实现该接口")
	fmt.Println()

	// 创建Rectangle实例
	rectangle := Rectangle{
		Width:  5.0,
		Height: 3.0,
	}
	fmt.Println("创建矩形实例:")
	printShapeInfo(rectangle)

	// 创建Circle实例
	circle := Circle{
		Radius: 4.0,
	}
	fmt.Println("创建圆形实例:")
	printShapeInfo(circle)

	// 演示多态性：使用接口类型
	fmt.Println("=== 多态性演示 ===")
	shapes := []Shape{rectangle, circle}
	
	fmt.Println("所有形状的信息:")
	for i, shape := range shapes {
		fmt.Printf("形状 %d:\n", i+1)
		printShapeInfo(shape)
	}

	// 计算总面积和总周长
	totalArea := calculateTotalArea(shapes)
	totalPerimeter := calculateTotalPerimeter(shapes)
	
	fmt.Printf("所有形状的总面积: %.2f\n", totalArea)
	fmt.Printf("所有形状的总周长: %.2f\n", totalPerimeter)
	fmt.Println()

	// 演示接口的灵活性
	fmt.Println("=== 接口灵活性演示 ===")
	
	// 创建不同尺寸的形状
	shapes2 := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
		Rectangle{Width: 3, Height: 8},
		Circle{Radius: 2.5},
	}
	
	fmt.Println("不同尺寸的形状:")
	for i, shape := range shapes2 {
		fmt.Printf("形状 %d: %v\n", i+1, shape)
		fmt.Printf("  面积: %.2f, 周长: %.2f\n", shape.Area(), shape.Perimeter())
	}
	
	fmt.Printf("\n所有形状的总面积: %.2f\n", calculateTotalArea(shapes2))
	fmt.Printf("所有形状的总周长: %.2f\n", calculateTotalPerimeter(shapes2))

	// 演示类型断言
	fmt.Println("\n=== 类型断言演示 ===")
	for i, shape := range shapes {
		fmt.Printf("形状 %d: ", i+1)
		
		// 类型断言
		if rect, ok := shape.(Rectangle); ok {
			fmt.Printf("这是一个矩形，宽: %.2f, 高: %.2f\n", rect.Width, rect.Height)
		} else if circ, ok := shape.(Circle); ok {
			fmt.Printf("这是一个圆形，半径: %.2f\n", circ.Radius)
		} else {
			fmt.Println("未知形状类型")
		}
	}
}
