package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印奇数的协程函数
func printOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知WaitGroup
	
	fmt.Println("奇数协程开始执行...")
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 添加小延迟以便观察并发执行
	}
	fmt.Println("奇数协程执行完毕")
}

// 打印偶数的协程函数
func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知WaitGroup
	
	fmt.Println("偶数协程开始执行...")
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 添加小延迟以便观察并发执行
	}
	fmt.Println("偶数协程执行完毕")
}

func main() {
	fmt.Println("=== Go协程并发执行示例 ===")
	fmt.Println("启动两个协程：一个打印奇数，一个打印偶数")
	fmt.Println()
	
	// 创建WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 等待2个协程完成
	
	// 使用go关键字启动第一个协程（打印奇数）
	go printOddNumbers(&wg)
	
	// 使用go关键字启动第二个协程（打印偶数）
	go printEvenNumbers(&wg)
	
	fmt.Println("主协程：已启动两个子协程，等待它们完成...")
	
	// 等待所有协程完成
	wg.Wait()
	
	fmt.Println()
	fmt.Println("主协程：所有子协程已完成！")
	fmt.Println("程序执行完毕")
}
