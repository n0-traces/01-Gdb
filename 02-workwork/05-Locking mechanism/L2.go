package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 使用原子操作的计数器结构体
type AtomicCounter struct {
	value int64
}

// 原子递增操作
func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.value, 1)
}

// 原子递减操作
func (ac *AtomicCounter) Decrement() {
	atomic.AddInt64(&ac.value, -1)
}

// 原子获取值
func (ac *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&ac.value)
}

// 原子设置值
func (ac *AtomicCounter) SetValue(newValue int64) {
	atomic.StoreInt64(&ac.value, newValue)
}

// 原子比较并交换操作
func (ac *AtomicCounter) CompareAndSwap(oldValue, newValue int64) bool {
	return atomic.CompareAndSwapInt64(&ac.value, oldValue, newValue)
}

// 原子重置操作
func (ac *AtomicCounter) Reset() {
	atomic.StoreInt64(&ac.value, 0)
}

// 使用原子操作的32位计数器
type AtomicCounter32 struct {
	value int32
}

func (ac *AtomicCounter32) Increment() {
	atomic.AddInt32(&ac.value, 1)
}

func (ac *AtomicCounter32) GetValue() int32 {
	return atomic.LoadInt32(&ac.value)
}

// 使用原子操作的指针计数器
type AtomicPointerCounter struct {
	value *int64
}

func NewAtomicPointerCounter() *AtomicPointerCounter {
	var value int64
	return &AtomicPointerCounter{
		value: &value,
	}
}

func (apc *AtomicPointerCounter) Increment() {
	atomic.AddInt64(apc.value, 1)
}

func (apc *AtomicPointerCounter) GetValue() int64 {
	return atomic.LoadInt64(apc.value)
}

// 工作协程：对原子计数器进行递增操作
func atomicWorker(counter *AtomicCounter, workerID int, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("原子工作协程 %d 开始工作...\n", workerID)
	
	for i := 0; i < iterations; i++ {
		counter.Increment()
		
		// 每100次操作显示一次进度
		if (i+1)%100 == 0 {
			fmt.Printf("原子工作协程 %d: 已完成 %d 次递增操作\n", workerID, i+1)
		}
	}
	
	fmt.Printf("原子工作协程 %d 完成工作\n", workerID)
}

// 32位原子工作协程
func atomicWorker32(counter *AtomicCounter32, workerID int, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("32位原子工作协程 %d 开始工作...\n", workerID)
	
	for i := 0; i < iterations; i++ {
		counter.Increment()
		
		if (i+1)%100 == 0 {
			fmt.Printf("32位原子工作协程 %d: 已完成 %d 次递增操作\n", workerID, i+1)
		}
	}
	
	fmt.Printf("32位原子工作协程 %d 完成工作\n", workerID)
}

// 演示原子操作的基本使用
func atomicCounterExample() {
	fmt.Println("=== 原子操作计数器示例 ===")
	
	// 创建原子计数器
	atomicCounter := &AtomicCounter{value: 0}
	
	// 设置参数
	workerCount := 10
	iterationsPerWorker := 1000
	expectedTotal := int64(workerCount * iterationsPerWorker)
	
	var wg sync.WaitGroup
	
	fmt.Printf("启动 %d 个原子工作协程，每个协程执行 %d 次递增操作\n", workerCount, iterationsPerWorker)
	fmt.Printf("期望的最终值: %d\n", expectedTotal)
	fmt.Println()
	
	startTime := time.Now()
	
	// 启动所有工作协程
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go atomicWorker(atomicCounter, i, iterationsPerWorker, &wg)
	}
	
	// 等待所有协程完成
	wg.Wait()
	
	duration := time.Since(startTime)
	finalValue := atomicCounter.GetValue()
	
	fmt.Println()
	fmt.Printf("所有原子工作协程已完成\n")
	fmt.Printf("最终计数值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", expectedTotal)
	fmt.Printf("是否一致: %t\n", finalValue == expectedTotal)
	fmt.Printf("总耗时: %v\n", duration)
	fmt.Printf("平均每次操作耗时: %v\n", duration/time.Duration(expectedTotal))
}

// 演示32位原子操作
func atomic32Example() {
	fmt.Println("\n=== 32位原子操作示例 ===")
	
	atomicCounter32 := &AtomicCounter32{value: 0}
	
	workerCount := 10
	iterationsPerWorker := 1000
	expectedTotal := int32(workerCount * iterationsPerWorker)
	
	var wg sync.WaitGroup
	
	fmt.Printf("启动 %d 个32位原子工作协程\n", workerCount)
	
	startTime := time.Now()
	
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go atomicWorker32(atomicCounter32, i, iterationsPerWorker, &wg)
	}
	
	wg.Wait()
	
	duration := time.Since(startTime)
	finalValue := atomicCounter32.GetValue()
	
	fmt.Printf("32位原子计数器最终值: %d (期望: %d)\n", finalValue, expectedTotal)
	fmt.Printf("32位原子操作耗时: %v\n", duration)
}

// 演示原子操作的比较和交换
func compareAndSwapExample() {
	fmt.Println("\n=== 原子比较和交换示例 ===")
	
	counter := &AtomicCounter{value: 0}
	
	// 演示CompareAndSwap操作
	fmt.Printf("初始值: %d\n", counter.GetValue())
	
	// 尝试将0替换为10
	success := counter.CompareAndSwap(0, 10)
	fmt.Printf("CompareAndSwap(0, 10): %t\n", success)
	fmt.Printf("当前值: %d\n", counter.GetValue())
	
	// 尝试将0替换为20（会失败，因为当前值是10）
	success = counter.CompareAndSwap(0, 20)
	fmt.Printf("CompareAndSwap(0, 20): %t\n", success)
	fmt.Printf("当前值: %d\n", counter.GetValue())
	
	// 尝试将10替换为20（会成功）
	success = counter.CompareAndSwap(10, 20)
	fmt.Printf("CompareAndSwap(10, 20): %t\n", success)
	fmt.Printf("当前值: %d\n", counter.GetValue())
}

// 演示原子操作的其他类型
func otherAtomicTypesExample() {
	fmt.Println("\n=== 其他原子操作类型示例 ===")
	
	// 原子布尔值
	var atomicBool int32
	fmt.Printf("原子布尔值初始状态: %t\n", atomic.LoadInt32(&atomicBool) != 0)
	
	atomic.StoreInt32(&atomicBool, 1)
	fmt.Printf("原子布尔值设置后: %t\n", atomic.LoadInt32(&atomicBool) != 0)
	
	// 原子指针
	var atomicPtr atomic.Value
	value := "Hello, Atomic!"
	atomicPtr.Store(value)
	
	loadedValue := atomicPtr.Load().(string)
	fmt.Printf("原子指针存储的值: %s\n", loadedValue)
	
	// 原子Uint64
	var atomicUint64 uint64
	atomic.AddUint64(&atomicUint64, 100)
	fmt.Printf("原子Uint64值: %d\n", atomic.LoadUint64(&atomicUint64))
}

// 互斥锁计数器（用于性能对比）
type MutexCounter struct {
	mu      sync.Mutex
	counter int64
}

func (mc *MutexCounter) Increment() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.counter++
}

func (mc *MutexCounter) GetValue() int64 {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	return mc.counter
}

// 性能对比：原子操作 vs 互斥锁
func performanceComparison() {
	fmt.Println("\n=== 性能对比：原子操作 vs 互斥锁 ===")
	
	workerCount := 10
	iterationsPerWorker := 1000
	
	// 测试原子操作
	atomicCounter := &AtomicCounter{value: 0}
	var wg1 sync.WaitGroup
	
	start1 := time.Now()
	for i := 0; i < workerCount; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for j := 0; j < iterationsPerWorker; j++ {
				atomicCounter.Increment()
			}
		}()
	}
	wg1.Wait()
	duration1 := time.Since(start1)
	
	// 测试互斥锁
	mutexCounter := &MutexCounter{counter: 0}
	var wg2 sync.WaitGroup
	
	start2 := time.Now()
	for i := 0; i < workerCount; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < iterationsPerWorker; j++ {
				mutexCounter.Increment()
			}
		}()
	}
	wg2.Wait()
	duration2 := time.Since(start2)
	
	fmt.Printf("原子操作: %v, 最终值: %d\n", duration1, atomicCounter.GetValue())
	fmt.Printf("互斥锁: %v, 最终值: %d\n", duration2, mutexCounter.GetValue())
	fmt.Printf("性能差异: %.2fx (原子操作更快)\n", float64(duration2)/float64(duration1))
}

// 演示原子操作的高级用法
func advancedAtomicExample() {
	fmt.Println("\n=== 高级原子操作示例 ===")
	
	// 原子操作实现自旋锁
	var spinLock int32
	
	// 尝试获取锁
	acquireLock := func() bool {
		return atomic.CompareAndSwapInt32(&spinLock, 0, 1)
	}
	
	// 释放锁
	releaseLock := func() {
		atomic.StoreInt32(&spinLock, 0)
	}
	
	// 演示自旋锁的使用
	fmt.Println("尝试获取自旋锁...")
	if acquireLock() {
		fmt.Println("成功获取锁")
		time.Sleep(100 * time.Millisecond)
		releaseLock()
		fmt.Println("释放锁")
	} else {
		fmt.Println("获取锁失败")
	}
	
	// 原子操作实现计数器数组
	var counters [5]int64
	
	// 并发更新不同的计数器
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counterIndex := j % 5
				atomic.AddInt64(&counters[counterIndex], 1)
			}
		}(i)
	}
	
	wg.Wait()
	
	fmt.Println("计数器数组最终值:")
	for i, value := range counters {
		fmt.Printf("  计数器[%d]: %d\n", i, value)
	}
}

func main() {
	fmt.Println("=== Go原子操作示例 ===")
	fmt.Println("使用sync/atomic包实现无锁计数器")
	fmt.Println()

	// 示例1：基本的原子操作计数器
	atomicCounterExample()

	// 示例2：32位原子操作
	atomic32Example()

	// 示例3：原子比较和交换
	compareAndSwapExample()

	// 示例4：其他原子操作类型
	otherAtomicTypesExample()

	// 示例5：性能对比
	performanceComparison()

	// 示例6：高级原子操作
	advancedAtomicExample()

	fmt.Println("\n=== 程序执行完毕 ===")
	fmt.Println("所有原子操作示例已完成")
}
