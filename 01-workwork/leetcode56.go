package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	current := intervals[0]
	result = append(result, current)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= current[1] {
			current[1] = max(current[1], intervals[i][1])
		} else {
			current = intervals[i]
			result = append(result, current)
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试 merge 函数
	fmt.Println("=== 测试 merge 函数 ===")

	// 测试用例1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("输入：intervals = %v\n", intervals1)
	result1 := merge(intervals1)
	fmt.Printf("输出：%v\n", result1)
	fmt.Printf("解释：区间 [1,3] 和 [2,6] 重叠，合并为 [1,6]\n\n")
}
