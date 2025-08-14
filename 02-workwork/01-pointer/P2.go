package main

import "fmt"

// 定义一个函数，接收一个整数切片的指针作为参数
// 该函数会将切片中的每个元素乘以2
func multiplySliceByTwo(slicePtr *[]int) {
	// 通过指针访问切片
	slice := *slicePtr
	
	// 遍历切片，将每个元素乘以2
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
	
	fmt.Printf("在函数内部：切片中的每个元素都已乘以2\n")
}

// 另一个实现方式：直接使用指针操作
func multiplySliceByTwoDirect(slicePtr *[]int) {
	// 直接通过指针操作切片元素
	for i := 0; i < len(*slicePtr); i++ {
		(*slicePtr)[i] *= 2
	}
	
	fmt.Printf("在函数内部（直接指针操作）：切片中的每个元素都已乘以2\n")
}

func main() {
	// 定义一个整数切片
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片：%v\n", numbers)
	
	// 调用函数，传递切片的地址（指针）
	multiplySliceByTwo(&numbers)
	
	// 输出修改后的切片
	fmt.Printf("修改后的切片：%v\n", numbers)
	
	// 演示另一种实现方式
	numbers2 := []int{10, 20, 30, 40, 50}
	fmt.Printf("\n原始切片2：%v\n", numbers2)
	
	multiplySliceByTwoDirect(&numbers2)
	fmt.Printf("修改后的切片2：%v\n", numbers2)
	
	// 演示指针的地址
	fmt.Printf("\n切片numbers的地址：%p\n", &numbers)
	fmt.Printf("切片numbers2的地址：%p\n", &numbers2)
}
