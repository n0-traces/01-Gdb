package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if top != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	// 测试 isValid 函数
	fmt.Println("=== 测试 isValid 函数 ===")

	// 测试用例1
	s1 := "()"
	result1 := isValid(s1)
	fmt.Printf("输入：s = \"%s\"\n", s1)
	fmt.Printf("输出：%t\n", result1)
	fmt.Printf("解释：括号匹配正确\n\n")

	// 测试用例2
	s2 := "()[]{}"
	result2 := isValid(s2)
	fmt.Printf("输入：s = \"%s\"\n", s2)
	fmt.Printf("输出：%t\n", result2)
	fmt.Printf("解释：所有括号都匹配正确\n\n")

	// 测试用例3
	s3 := "(]"
	result3 := isValid(s3)
	fmt.Printf("输入：s = \"%s\"\n", s3)
	fmt.Printf("输出：%t\n", result3)
	fmt.Printf("解释：括号类型不匹配\n\n")

	// 测试用例4
	s4 := "([])"
	result4 := isValid(s4)
	fmt.Printf("输入：s = \"%s\"\n", s4)
	fmt.Printf("输出：%t\n", result4)
	fmt.Printf("解释：括号匹配正确\n\n")

	// 测试用例5
	s5 := "([)]"
	result5 := isValid(s5)
	fmt.Printf("输入：s = \"%s\"\n", s5)
	fmt.Printf("输出：%t\n", result5)
	fmt.Printf("解释：括号顺序不正确\n\n")

	// 额外测试用例
	testCases := []string{
		"",
		"(",
		")",
		"(((",
		")))",
		"({[]})",
		"({[}])",
		"((()))",
		"(()",
		"())",
	}

	fmt.Println("=== 额外测试用例 ===")
	for i, s := range testCases {
		result := isValid(s)
		fmt.Printf("测试用例 %d: \"%s\" → %t\n", i+1, s, result)
	}
}
