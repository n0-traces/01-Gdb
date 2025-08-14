package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者协程：向通道发送数据
func producer(ch chan int, count int, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	defer close(ch) // 发送完毕后关闭通道
	
	fmt.Printf("%s: 开始发送 %d 个数据...\n", name, count)
	startTime := time.Now()
	
	for i := 1; i <= count; i++ {
		ch <- i // 发送数据到通道
		if i%20 == 0 {
			fmt.Printf("%s: 已发送 %d 个数据\n", name, i)
		}
		time.Sleep(10 * time.Millisecond) // 模拟发送延迟
	}
	
	duration := time.Since(startTime)
	fmt.Printf("%s: 发送完成，耗时: %v\n", name, duration)
}

// 消费者协程：从通道接收数据
func consumer(ch chan int, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	
	fmt.Printf("%s: 开始接收数据...\n", name)
	startTime := time.Now()
	
	count := 0
	for num := range ch {
		count++
		if count%20 == 0 {
			fmt.Printf("%s: 已接收 %d 个数据，当前数据: %d\n", name, count, num)
		}
		time.Sleep(15 * time.Millisecond) // 模拟处理延迟
	}
	
	duration := time.Since(startTime)
	fmt.Printf("%s: 接收完成，共接收 %d 个数据，耗时: %v\n", name, count, duration)
}

// 带缓冲通道示例
func bufferedChannelExample(bufferSize int, dataCount int) {
	fmt.Printf("\n=== 带缓冲通道示例 (缓冲区大小: %d) ===\n", bufferSize)
	
	// 创建带缓冲的通道
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// 启动消费者协程
	go consumer(ch, &wg, "消费者")
	
	// 启动生产者协程
	go producer(ch, dataCount, &wg, "生产者")
	
	// 等待所有协程完成
	wg.Wait()
	
	fmt.Printf("缓冲区大小 %d 的示例完成\n", bufferSize)
}

// 无缓冲通道示例（对比）
func unbufferedChannelExample(dataCount int) {
	fmt.Printf("\n=== 无缓冲通道示例 ===\n")
	
	// 创建无缓冲通道
	ch := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	// 启动消费者协程
	go consumer(ch, &wg, "消费者(无缓冲)")
	
	// 启动生产者协程
	go producer(ch, dataCount, &wg, "生产者(无缓冲)")
	
	// 等待所有协程完成
	wg.Wait()
	
	fmt.Println("无缓冲通道示例完成")
}

// 性能对比测试
func performanceComparison() {
	fmt.Println("\n=== 性能对比测试 ===")
	
	dataCount := 100
	bufferSizes := []int{0, 1, 5, 10, 20, 50}
	
	for _, bufferSize := range bufferSizes {
		fmt.Printf("\n--- 测试缓冲区大小: %d ---\n", bufferSize)
		
		startTime := time.Now()
		
		if bufferSize == 0 {
			// 无缓冲通道
			ch := make(chan int)
			var wg sync.WaitGroup
			wg.Add(2)
			
			go consumer(ch, &wg, "消费者")
			go producer(ch, dataCount, &wg, "生产者")
			wg.Wait()
		} else {
			// 带缓冲通道
			ch := make(chan int, bufferSize)
			var wg sync.WaitGroup
			wg.Add(2)
			
			go consumer(ch, &wg, "消费者")
			go producer(ch, dataCount, &wg, "生产者")
			wg.Wait()
		}
		
		duration := time.Since(startTime)
		fmt.Printf("缓冲区大小 %d: 总耗时 %v\n", bufferSize, duration)
	}
}

// 多生产者多消费者示例
func multiProducerConsumerExample() {
	fmt.Println("\n=== 多生产者多消费者示例 ===")
	
	bufferSize := 20
	producerCount := 3
	consumerCount := 2
	dataPerProducer := 30
	
	// 创建带缓冲的通道
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup
	
	// 启动多个消费者
	for i := 1; i <= consumerCount; i++ {
		wg.Add(1)
		go consumer(ch, &wg, fmt.Sprintf("消费者-%d", i))
	}
	
	// 启动多个生产者
	for i := 1; i <= producerCount; i++ {
		wg.Add(1)
		go producer(ch, dataPerProducer, &wg, fmt.Sprintf("生产者-%d", i))
	}
	
	// 等待所有协程完成
	wg.Wait()
	
	fmt.Printf("多生产者多消费者示例完成 (生产者: %d, 消费者: %d, 总数据: %d)\n", 
		producerCount, consumerCount, producerCount*dataPerProducer)
}

// 通道容量监控示例
func channelCapacityMonitor() {
	fmt.Println("\n=== 通道容量监控示例 ===")
	
	bufferSize := 10
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup
	
	wg.Add(3)
	
	// 监控协程
	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("监控: 通道长度 %d/%d\n", len(ch), cap(ch))
		}
	}()
	
	// 快速生产者
	go func() {
		defer wg.Done()
		defer close(ch)
		
		for i := 1; i <= 50; i++ {
			ch <- i
			time.Sleep(20 * time.Millisecond)
		}
	}()
	
	// 慢速消费者
	go func() {
		defer wg.Done()
		
		for num := range ch {
			fmt.Printf("消费者: 接收 %d\n", num)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Println("通道容量监控示例完成")
}

func main() {
	fmt.Println("=== Go带缓冲通道示例 ===")
	fmt.Println("展示通道缓冲机制的工作原理")
	fmt.Println()

	// 示例1：不同缓冲大小的对比
	bufferedChannelExample(5, 100)   // 缓冲区大小5
	bufferedChannelExample(20, 100)  // 缓冲区大小20
	bufferedChannelExample(50, 100)  // 缓冲区大小50

	// 示例2：无缓冲通道对比
	unbufferedChannelExample(100)

	// 示例3：性能对比测试
	performanceComparison()

	// 示例4：多生产者多消费者
	multiProducerConsumerExample()

	// 示例5：通道容量监控
	channelCapacityMonitor()

	fmt.Println("\n=== 程序执行完毕 ===")
	fmt.Println("所有带缓冲通道示例已完成")
}
