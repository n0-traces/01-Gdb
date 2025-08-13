package main

import "fmt"

// singleNumber 函数：找出只出现一次的元素
func singleNumber(nums []int) int {
	// 使用异或运算的特性
	// 任何数与0异或都是它本身
	// 任何数与它本身异或都是0
	// 异或运算满足交换律和结合律
	result := 0

	// 遍历数组，将所有数字进行异或运算
	for _, num := range nums {
		result ^= num
	}

	return result
}

func main() {
	// 测试 singleNumber 函数
	fmt.Println("=== 测试 singleNumber 函数 ===")

	// 测试用例1
	nums1 := []int{2, 2, 1}
	fmt.Printf("输入：nums = %v\n", nums1)
	result1 := singleNumber(nums1)
	fmt.Printf("输出：%d\n", result1)
	fmt.Printf("解释：数字 1 只出现了一次\n\n")

	// 测试用例2
	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Printf("输入：nums = %v\n", nums2)
	result2 := singleNumber(nums2)
	fmt.Printf("输出：%d\n", result2)
	fmt.Printf("解释：数字 4 只出现了一次\n\n")

	// 测试用例3
	nums3 := []int{1}
	fmt.Printf("输入：nums = %v\n", nums3)
	result3 := singleNumber(nums3)
	fmt.Printf("输出：%d\n", result3)
	fmt.Printf("解释：数字 1 只出现了一次\n\n")

	// 额外测试用例
	testCases := [][]int{
		{1, 1, 2, 2, 3},
		{5, 5, 6, 6, 7, 7, 8},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 0},
		{-1, -1, -2},
	}

	fmt.Println("=== 额外测试用例 ===")
	for i, nums := range testCases {
		result := singleNumber(nums)
		fmt.Printf("测试用例 %d: %v → %d\n", i+1, nums, result)
	}

	// 演示异或运算的特性
	fmt.Println("\n=== 异或运算特性演示 ===")
	fmt.Println("异或运算的特性：")
	fmt.Println("1. a ^ 0 = a")
	fmt.Println("2. a ^ a = 0")
	fmt.Println("3. a ^ b ^ a = b (交换律和结合律)")

	// 具体例子演示
	fmt.Println("\n具体例子：")
	example := []int{4, 1, 2, 1, 2}
	fmt.Printf("数组：%v\n", example)
	fmt.Println("异或过程：")
	result := 0
	for i, num := range example {
		result ^= num
		fmt.Printf("  步骤 %d: %d ^ %d = %d\n", i+1, result^num, num, result)
	}
	fmt.Printf("最终结果：%d\n", result)
}
