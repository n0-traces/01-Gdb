package main

import (
	"fmt"
	"sync"
	"time"
)

// 带互斥锁的计数器结构体
type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

// 安全的递增方法
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()   // 获取锁
	defer sc.mu.Unlock() // 确保锁会被释放
	
	sc.counter++
}

// 安全的获取计数值
func (sc *SafeCounter) GetValue() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	
	return sc.counter
}

// 安全的递减方法
func (sc *SafeCounter) Decrement() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	
	sc.counter--
}

// 安全的重置方法
func (sc *SafeCounter) Reset() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	
	sc.counter = 0
}

// 不安全的计数器（用于对比）
type UnsafeCounter struct {
	counter int
}

func (uc *UnsafeCounter) Increment() {
	uc.counter++ // 没有锁保护，会发生数据竞争
}

func (uc *UnsafeCounter) GetValue() int {
	return uc.counter
}

// 工作协程：对计数器进行递增操作
func worker(counter *SafeCounter, workerID int, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("工作协程 %d 开始工作...\n", workerID)
	
	for i := 0; i < iterations; i++ {
		counter.Increment()
		
		// 每100次操作显示一次进度
		if (i+1)%100 == 0 {
			fmt.Printf("工作协程 %d: 已完成 %d 次递增操作\n", workerID, i+1)
		}
	}
	
	fmt.Printf("工作协程 %d 完成工作\n", workerID)
}

// 不安全的工作协程（用于对比）
func unsafeWorker(counter *UnsafeCounter, workerID int, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("不安全工作协程 %d 开始工作...\n", workerID)
	
	for i := 0; i < iterations; i++ {
		counter.Increment()
		
		if (i+1)%100 == 0 {
			fmt.Printf("不安全工作协程 %d: 已完成 %d 次递增操作\n", workerID, i+1)
		}
	}
	
	fmt.Printf("不安全工作协程 %d 完成工作\n", workerID)
}

// 演示安全的计数器使用
func safeCounterExample() {
	fmt.Println("=== 安全的计数器示例 ===")
	
	// 创建安全的计数器
	safeCounter := &SafeCounter{counter: 0}
	
	// 设置参数
	workerCount := 10
	iterationsPerWorker := 1000
	expectedTotal := workerCount * iterationsPerWorker
	
	var wg sync.WaitGroup
	
	fmt.Printf("启动 %d 个工作协程，每个协程执行 %d 次递增操作\n", workerCount, iterationsPerWorker)
	fmt.Printf("期望的最终值: %d\n", expectedTotal)
	fmt.Println()
	
	startTime := time.Now()
	
	// 启动所有工作协程
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(safeCounter, i, iterationsPerWorker, &wg)
	}
	
	// 等待所有协程完成
	wg.Wait()
	
	duration := time.Since(startTime)
	finalValue := safeCounter.GetValue()
	
	fmt.Println()
	fmt.Printf("所有工作协程已完成\n")
	fmt.Printf("最终计数值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", expectedTotal)
	fmt.Printf("是否一致: %t\n", finalValue == expectedTotal)
	fmt.Printf("总耗时: %v\n", duration)
	fmt.Printf("平均每次操作耗时: %v\n", duration/time.Duration(expectedTotal))
}

// 演示不安全的计数器（数据竞争）
func unsafeCounterExample() {
	fmt.Println("\n=== 不安全的计数器示例（数据竞争） ===")
	
	// 创建不安全的计数器
	unsafeCounter := &UnsafeCounter{counter: 0}
	
	// 设置参数
	workerCount := 10
	iterationsPerWorker := 1000
	expectedTotal := workerCount * iterationsPerWorker
	
	var wg sync.WaitGroup
	
	fmt.Printf("启动 %d 个不安全工作协程，每个协程执行 %d 次递增操作\n", workerCount, iterationsPerWorker)
	fmt.Printf("期望的最终值: %d\n", expectedTotal)
	fmt.Println("注意：由于数据竞争，实际结果可能小于期望值")
	fmt.Println()
	
	startTime := time.Now()
	
	// 启动所有工作协程
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go unsafeWorker(unsafeCounter, i, iterationsPerWorker, &wg)
	}
	
	// 等待所有协程完成
	wg.Wait()
	
	duration := time.Since(startTime)
	finalValue := unsafeCounter.GetValue()
	
	fmt.Println()
	fmt.Printf("所有不安全工作协程已完成\n")
	fmt.Printf("最终计数值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", expectedTotal)
	fmt.Printf("是否一致: %t\n", finalValue == expectedTotal)
	fmt.Printf("丢失的操作数: %d\n", expectedTotal-finalValue)
	fmt.Printf("总耗时: %v\n", duration)
	fmt.Printf("平均每次操作耗时: %v\n", duration/time.Duration(expectedTotal))
}

// 演示读写锁的使用
func rwMutexExample() {
	fmt.Println("\n=== 读写锁示例 ===")
	
	type SafeData struct {
		mu    sync.RWMutex
		data  map[string]int
	}
	
	safeData := &SafeData{
		data: make(map[string]int),
	}
	
	// 写操作
	writeWorker := func(key string, value int, wg *sync.WaitGroup) {
		defer wg.Done()
		
		safeData.mu.Lock() // 写锁
		defer safeData.mu.Unlock()
		
		safeData.data[key] = value
		fmt.Printf("写入: %s = %d\n", key, value)
		time.Sleep(10 * time.Millisecond)
	}
	
	// 读操作
	readWorker := func(key string, wg *sync.WaitGroup) {
		defer wg.Done()
		
		safeData.mu.RLock() // 读锁
		defer safeData.mu.RUnlock()
		
		value, exists := safeData.data[key]
		if exists {
			fmt.Printf("读取: %s = %d\n", key, value)
		} else {
			fmt.Printf("读取: %s 不存在\n", key)
		}
		time.Sleep(5 * time.Millisecond)
	}
	
	var wg sync.WaitGroup
	
	// 启动写操作
	wg.Add(3)
	go writeWorker("A", 1, &wg)
	go writeWorker("B", 2, &wg)
	go writeWorker("C", 3, &wg)
	
	// 启动读操作
	wg.Add(6)
	go readWorker("A", &wg)
	go readWorker("B", &wg)
	go readWorker("C", &wg)
	go readWorker("A", &wg)
	go readWorker("B", &wg)
	go readWorker("C", &wg)
	
	wg.Wait()
	
	fmt.Println("读写锁示例完成")
}

// 性能对比测试
func performanceComparison() {
	fmt.Println("\n=== 性能对比测试 ===")
	
	workerCount := 10
	iterationsPerWorker := 1000
	
	// 测试安全计数器
	safeCounter := &SafeCounter{counter: 0}
	var wg1 sync.WaitGroup
	
	start1 := time.Now()
	for i := 0; i < workerCount; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for j := 0; j < iterationsPerWorker; j++ {
				safeCounter.Increment()
			}
		}()
	}
	wg1.Wait()
	duration1 := time.Since(start1)
	
	// 测试不安全计数器
	unsafeCounter := &UnsafeCounter{counter: 0}
	var wg2 sync.WaitGroup
	
	start2 := time.Now()
	for i := 0; i < workerCount; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < iterationsPerWorker; j++ {
				unsafeCounter.Increment()
			}
		}()
	}
	wg2.Wait()
	duration2 := time.Since(start2)
	
	fmt.Printf("安全计数器（带锁）: %v, 最终值: %d\n", duration1, safeCounter.GetValue())
	fmt.Printf("不安全计数器（无锁）: %v, 最终值: %d\n", duration2, unsafeCounter.GetValue())
	fmt.Printf("性能差异: %.2fx\n", float64(duration1)/float64(duration2))
}

func main() {
	fmt.Println("=== Go互斥锁示例 ===")
	fmt.Println("使用sync.Mutex保护共享计数器")
	fmt.Println()

	// 示例1：安全的计数器
	safeCounterExample()

	// 示例2：不安全的计数器（数据竞争）
	unsafeCounterExample()

	// 示例3：读写锁示例
	rwMutexExample()

	// 示例4：性能对比
	performanceComparison()

	fmt.Println("\n=== 程序执行完毕 ===")
	fmt.Println("所有互斥锁示例已完成")
}
