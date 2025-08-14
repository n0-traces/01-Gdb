package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务结构体
type Task struct {
	ID       int
	Name     string
	Function func() interface{} // 任务函数
}

// TaskResult 定义任务执行结果
type TaskResult struct {
	TaskID      int
	TaskName    string
	Result      interface{}
	Duration    time.Duration
	StartTime   time.Time
	EndTime     time.Time
	Error       error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks    []Task
	results  []TaskResult
	mutex    sync.RWMutex
	wg       sync.WaitGroup
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务到调度器
func (ts *TaskScheduler) AddTask(name string, taskFunc func() interface{}) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	
	task := Task{
		ID:       len(ts.tasks) + 1,
		Name:     name,
		Function: taskFunc,
	}
	ts.tasks = append(ts.tasks, task)
}

// executeTask 执行单个任务并记录结果
func (ts *TaskScheduler) executeTask(task Task) {
	defer ts.wg.Done()
	
	startTime := time.Now()
	
	// 执行任务
	result := task.Function()
	
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	
	// 记录结果
	taskResult := TaskResult{
		TaskID:    task.ID,
		TaskName:  task.Name,
		Result:    result,
		Duration:  duration,
		StartTime: startTime,
		EndTime:   endTime,
	}
	
	// 线程安全地添加结果
	ts.mutex.Lock()
	ts.results = append(ts.results, taskResult)
	ts.mutex.Unlock()
	
	fmt.Printf("任务 %d (%s) 执行完成，耗时: %v\n", task.ID, task.Name, duration)
}

// RunAll 并发执行所有任务
func (ts *TaskScheduler) RunAll() {
	fmt.Printf("开始执行 %d 个任务...\n", len(ts.tasks))
	startTime := time.Now()
	
	// 为每个任务启动一个协程
	ts.wg.Add(len(ts.tasks))
	for _, task := range ts.tasks {
		go ts.executeTask(task)
	}
	
	// 等待所有任务完成
	ts.wg.Wait()
	
	totalDuration := time.Since(startTime)
	fmt.Printf("\n所有任务执行完成，总耗时: %v\n", totalDuration)
}

// GetResults 获取所有任务的执行结果
func (ts *TaskScheduler) GetResults() []TaskResult {
	ts.mutex.RLock()
	defer ts.mutex.RUnlock()
	
	results := make([]TaskResult, len(ts.results))
	copy(results, ts.results)
	return results
}

// PrintSummary 打印执行摘要
func (ts *TaskScheduler) PrintSummary() {
	results := ts.GetResults()
	
	fmt.Println("\n=== 任务执行摘要 ===")
	fmt.Printf("%-5s %-15s %-15s %-20s\n", "ID", "任务名", "执行时间", "结果")
	fmt.Println("------------------------------------------------------------")
	
	var totalDuration time.Duration
	for _, result := range results {
		fmt.Printf("%-5d %-15s %-15v %-20v\n", 
			result.TaskID, 
			result.TaskName, 
			result.Duration, 
			result.Result)
		totalDuration += result.Duration
	}
	
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("总执行时间: %v\n", totalDuration)
	fmt.Printf("并发执行节省时间: %v\n", totalDuration-time.Duration(len(results))*100*time.Millisecond)
}

// 示例任务函数
func task1() interface{} {
	time.Sleep(200 * time.Millisecond)
	return "任务1完成"
}

func task2() interface{} {
	time.Sleep(150 * time.Millisecond)
	return 42
}

func task3() interface{} {
	time.Sleep(300 * time.Millisecond)
	return []string{"a", "b", "c"}
}

func task4() interface{} {
	time.Sleep(100 * time.Millisecond)
	return map[string]int{"count": 100}
}

func task5() interface{} {
	time.Sleep(250 * time.Millisecond)
	return 3.14159
}

func main() {
	fmt.Println("=== Go协程任务调度器示例 ===")
	
	// 创建任务调度器
	scheduler := NewTaskScheduler()
	
	// 添加任务
	scheduler.AddTask("数据处理", task1)
	scheduler.AddTask("数值计算", task2)
	scheduler.AddTask("列表生成", task3)
	scheduler.AddTask("映射创建", task4)
	scheduler.AddTask("浮点运算", task5)
	
	// 并发执行所有任务
	scheduler.RunAll()
	
	// 打印执行摘要
	scheduler.PrintSummary()
	
	// 演示获取详细结果
	fmt.Println("\n=== 详细结果 ===")
	results := scheduler.GetResults()
	for _, result := range results {
		fmt.Printf("任务 %d (%s):\n", result.TaskID, result.TaskName)
		fmt.Printf("  开始时间: %v\n", result.StartTime.Format("15:04:05.000"))
		fmt.Printf("  结束时间: %v\n", result.EndTime.Format("15:04:05.000"))
		fmt.Printf("  执行时间: %v\n", result.Duration)
		fmt.Printf("  返回结果: %v\n", result.Result)
		fmt.Println()
	}
}
