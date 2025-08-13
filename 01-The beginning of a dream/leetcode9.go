package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x = x / 10
	}
	return x == reversed || x == reversed/10
}

func main() {
	// 测试 isPalindrome 函数
	fmt.Println("=== 测试 isPalindrome 函数 ===")

	x1 := 121
	result1 := isPalindrome(x1)
	fmt.Printf("输入：x = %d\n", x1)
	fmt.Printf("输出：%t\n", result1)
	fmt.Printf("解释：%d 是回文\n", x1)

	x2 := -121
	result2 := isPalindrome(x2)
	fmt.Printf("输入：x = %d\n", x2)
	fmt.Printf("输出：%t\n", result2)
	fmt.Printf("解释：%d 不是回文数（负数）\n\n", x2)

	x3 := 10
	result3 := isPalindrome(x3)
	fmt.Printf("输入：x = %d\n", x3)
	fmt.Printf("输出：%t\n", result3)
	fmt.Printf("解释：%d 不是回文数（末尾为0）\n\n", x3)

	testCases := []int{12321, 12345, 0, 1, 123, 1221, 1234}
	for _, x := range testCases {
		result := isPalindrome(x)
		fmt.Printf("输入：x = %d, 输出：%t\n", x, result)
	}

}
