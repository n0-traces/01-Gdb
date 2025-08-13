package main

import "fmt"

// plusOne 函数：对大整数数组加1
func plusOne(digits []int) []int {
	// 从右到左遍历数组
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果当前位不是9，直接加1并返回
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// 如果当前位是9，进位，当前位变为0
		digits[i] = 0
	}

	// 如果所有位都是9，需要在最前面添加一个1
	// 创建一个新的数组，长度为原数组长度+1
	result := make([]int, len(digits)+1)
	result[0] = 1
	// 其余位都是0（因为原数组所有位都进位了）

	return result
}

func main() {
	// 测试 plusOne 函数
	fmt.Println("=== 测试 plusOne 函数 ===")

	// 测试用例1
	digits1 := []int{1, 2, 3}
	fmt.Printf("输入：digits = %v\n", digits1)
	result1 := plusOne(digits1)
	fmt.Printf("输出：%v\n", result1)
	fmt.Printf("解释：123 + 1 = 124\n\n")

	// 测试用例2
	digits2 := []int{4, 3, 2, 1}
	fmt.Printf("输入：digits = %v\n", digits2)
	result2 := plusOne(digits2)
	fmt.Printf("输出：%v\n", result2)
	fmt.Printf("解释：4321 + 1 = 4322\n\n")

	// 测试用例3
	digits3 := []int{9}
	fmt.Printf("输入：digits = %v\n", digits3)
	result3 := plusOne(digits3)
	fmt.Printf("输出：%v\n", result3)
	fmt.Printf("解释：9 + 1 = 10\n\n")

	// 额外测试用例
	testCases := [][]int{
		{1, 2, 9},
		{9, 9, 9},
		{0},
		{1, 0, 0, 0},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}

	fmt.Println("=== 额外测试用例 ===")
	for i, digits := range testCases {
		// 创建副本避免影响其他测试
		digitsCopy := make([]int, len(digits))
		copy(digitsCopy, digits)

		result := plusOne(digitsCopy)
		fmt.Printf("测试用例 %d: %v → %v\n", i+1, digits, result)
	}
}
