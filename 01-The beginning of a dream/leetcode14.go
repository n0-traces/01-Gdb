package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for i := 0; i < minLen; i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:minLen]
}

func main() {
	// 测试 longestCommonPrefix 函数
	fmt.Println("=== 测试 longestCommonPrefix 函数 ===")

	strs1 := []string{"flower", "flow", "flight"}
	result1 := longestCommonPrefix(strs1)
	fmt.Printf("输入：strs = %v\n", strs1)
	fmt.Printf("输出：\"%s\"\n", result1)
	fmt.Printf("解释：最长公共前缀是 \"%s\"\n\n", result1)

	strs2 := []string{"dog", "racecar", "car"}
	result2 := longestCommonPrefix(strs2)
	fmt.Printf("输入：strs = %v\n", strs2)
	fmt.Printf("输出：\"%s\"\n", result2)
	fmt.Printf("解释：不存在公共前缀\n\n")

	testCases := [][]string{
		{"interspecies", "interstellar", "interstate"},
		{"throne", "throne"},
		{"throne", "dungeon"},
		{"", "b"},
		{"a"},
		{},
	}

	for i, strs := range testCases {
		result := longestCommonPrefix(strs)
		fmt.Printf("测试用例 %d: %v → \"%s\"\n", i+1, strs, result)
	}
	
}
