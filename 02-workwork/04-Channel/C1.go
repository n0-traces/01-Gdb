package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者协程：生成数字并发送到通道
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // 发送完毕后关闭通道
	
	fmt.Println("生产者协程开始工作...")
	
	for i := 1; i <= 10; i++ {
		fmt.Printf("生产者: 发送数字 %d\n", i)
		ch <- i // 发送数据到通道
		time.Sleep(100 * time.Millisecond) // 模拟工作延迟
	}
	
	fmt.Println("生产者协程: 所有数字已发送完毕")
}

// 消费者协程：从通道接收数字并打印
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Println("消费者协程开始工作...")
	
	// 方法1：使用for range循环接收数据
	for num := range ch {
		fmt.Printf("消费者: 接收到数字 %d\n", num)
		time.Sleep(150 * time.Millisecond) // 模拟处理延迟
	}
	
	fmt.Println("消费者协程: 所有数字已处理完毕")
}

// 消费者协程（方法2）：使用for循环和select
func consumerWithSelect(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Println("消费者协程(Select方式)开始工作...")
	
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				// 通道已关闭
				fmt.Println("消费者协程(Select): 通道已关闭，退出")
				return
			}
			fmt.Printf("消费者(Select): 接收到数字 %d\n", num)
			time.Sleep(150 * time.Millisecond)
		case <-time.After(1 * time.Second):
			// 超时处理（这里不会触发，因为数据会及时到达）
			fmt.Println("消费者协程(Select): 等待超时")
			return
		}
	}
}

// 双向通信示例：计算器协程
func calculator(inputCh chan int, outputCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(outputCh)
	
	fmt.Println("计算器协程开始工作...")
	
	for num := range inputCh {
		result := num * num // 计算平方
		fmt.Printf("计算器: 计算 %d 的平方 = %d\n", num, result)
		outputCh <- result
		time.Sleep(200 * time.Millisecond)
	}
	
	fmt.Println("计算器协程: 计算完成")
}

// 结果收集器协程
func resultCollector(outputCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Println("结果收集器协程开始工作...")
	
	var results []int
	for result := range outputCh {
		results = append(results, result)
		fmt.Printf("收集器: 收集到结果 %d\n", result)
	}
	
	fmt.Printf("结果收集器: 所有结果已收集完毕，共 %d 个结果\n", len(results))
	fmt.Printf("收集到的结果: %v\n", results)
}

// 带缓冲的通道示例
func bufferedChannelExample() {
	fmt.Println("\n=== 带缓冲通道示例 ===")
	
	// 创建带缓冲的通道，容量为3
	bufferedCh := make(chan int, 3)
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// 生产者协程
	go func() {
		defer wg.Done()
		defer close(bufferedCh)
		
		fmt.Println("缓冲通道生产者开始...")
		for i := 1; i <= 5; i++ {
			fmt.Printf("缓冲生产者: 发送 %d\n", i)
			bufferedCh <- i
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("缓冲通道生产者完成")
	}()
	
	// 消费者协程
	go func() {
		defer wg.Done()
		
		fmt.Println("缓冲通道消费者开始...")
		for num := range bufferedCh {
			fmt.Printf("缓冲消费者: 接收 %d\n", num)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("缓冲通道消费者完成")
	}()
	
	wg.Wait()
}

func main() {
	fmt.Println("=== Go通道通信示例 ===")
	fmt.Println("使用通道实现协程间通信")
	fmt.Println()

	// 示例1：基本的通道通信
	fmt.Println("=== 基本通道通信示例 ===")
	
	// 创建无缓冲通道
	ch := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// 启动生产者协程
	go producer(ch, &wg)
	
	// 启动消费者协程
	go consumer(ch, &wg)
	
	// 等待所有协程完成
	wg.Wait()
	
	fmt.Println("基本通道通信示例完成")
	fmt.Println()

	// 示例2：双向通信（计算器示例）
	fmt.Println("=== 双向通道通信示例 ===")
	
	inputCh := make(chan int)
	outputCh := make(chan int)
	
	wg.Add(3)
	
	// 启动计算器协程
	go calculator(inputCh, outputCh, &wg)
	
	// 启动结果收集器协程
	go resultCollector(outputCh, &wg)
	
	// 主协程发送数据
	go func() {
		defer wg.Done()
		defer close(inputCh)
		
		fmt.Println("主协程: 开始发送数据给计算器")
		for i := 1; i <= 5; i++ {
			fmt.Printf("主协程: 发送 %d 给计算器\n", i)
			inputCh <- i
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("主协程: 数据发送完毕")
	}()
	
	wg.Wait()
	fmt.Println("双向通道通信示例完成")
	
	// 示例3：带缓冲的通道
	bufferedChannelExample()
	
	fmt.Println("\n=== 程序执行完毕 ===")
	fmt.Println("所有通道通信示例已完成")
}
