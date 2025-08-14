package main

import "fmt"

// 定义一个函数，接收一个整数指针作为参数
// 该函数会将指针指向的值增加10
func increaseByTen(ptr *int) {
	// 通过指针修改原值
	*ptr += 10
	fmt.Printf("在函数内部：指针指向的值增加了10，当前值为：%d\n", *ptr)
}

func main() {
	// 定义一个整数变量
	num := 25
	fmt.Printf("调用函数前：num = %d\n", num)
	
	// 调用函数，传递num的地址（指针）
	increaseByTen(&num)
	
	// 输出修改后的值
	fmt.Printf("调用函数后：num = %d\n", num)
	
	// 演示指针的地址
	fmt.Printf("num的地址：%p\n", &num)
}
