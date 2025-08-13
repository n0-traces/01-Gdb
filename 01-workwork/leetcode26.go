package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	// 测试 removeDuplicates 函数
	fmt.Println("=== 测试 removeDuplicates 函数 ===")

	nums1 := []int{1, 1, 2}
	fmt.Printf("输入：nums = %v\n", nums1)
	k1 := removeDuplicates(nums1)
	fmt.Printf("输出：%d, nums = %v\n", k1, nums1[:k1])
	fmt.Printf("解释：函数返回新的长度 %d，前 %d 个元素为 %v\n\n", k1, k1, nums1[:k1])

	// 测试用例2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("输入：nums = %v\n", nums2)
	k2 := removeDuplicates(nums2)
	fmt.Printf("输出：%d, nums = %v\n", k2, nums2[:k2])
	fmt.Printf("解释：函数返回新的长度 %d，前 %d 个元素为 %v\n\n", k2, k2, nums2[:k2])

	// 额外测试用例
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1, 1},
		{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
		{},
		{1},
		{1, 1},
		{1, 2},
	}

	fmt.Println("=== 额外测试用例 ===")
	for i, nums := range testCases {
		// 创建副本避免影响其他测试
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)

		k := removeDuplicates(numsCopy)
		fmt.Printf("测试用例 %d: %v → %d, %v\n", i+1, nums, k, numsCopy[:k])
	}
}
