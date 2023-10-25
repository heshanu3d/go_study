package third

import (
	"context"
	"fmt"
	"math/rand"
	"time"
    "runtime"
)

// 异步任务提前完成
// PS D:\code\go\study> .\study.exe
// main thread set timeout 3s
// main thread select
// [ runtime.goexit ]  start
// [ runtime.goexit ]  select
// Task completed with result: 19

// 上下文任务超时
// PS D:\code\go\study> .\study.exe
// main thread set timeout 3s
// main thread select
// [ runtime.goexit ]  start
// main thread Task timed out


func Ctx() {
    fmt.Println("main thread start")
	// 创建一个根上下文
	rootCtx := context.Background()

	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(rootCtx, 3 * time.Second)
	defer cancel()
    fmt.Println("main thread set timeout 3s")

	// 创建一个用于接收结果的通道
	resultChan := make(chan int)

    fmt.Println("main thread set up a new task")
	// 启动一个并发任务
	go func(ctx context.Context) {
		// 模拟一些耗时的操作
        pc, _, _, _ := runtime.Caller(1)
        currentFunction := runtime.FuncForPC(pc).Name()
        fmt.Println("[", currentFunction, "] ", "start")
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

        fmt.Println("[", currentFunction, "] ", "select")
		// 检查上下文是否已被取消
		select {
		case <-ctx.Done():
			// fmt.Println("Task canceled")
            fmt.Println("[", currentFunction, "] ", "Task canceled")
			return
		default:
			// 如果上下文未被取消，则继续执行任务
			result := rand.Intn(100)
			resultChan <- result
		}
	}(ctx)

    fmt.Println("main thread select")
	// 等待任务完成或超时
	select {
	case result := <-resultChan:
		fmt.Printf("Task completed with result: %d\n", result)
	case <-ctx.Done():
        fmt.Println("main thread Task timed out")
	}
    time.Sleep(3 * time.Second)
    fmt.Println("main thread exit")
}

func Ctx_cancel() {
    fmt.Println("main thread start")
    // 创建一个带有取消功能的上下文
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

    fmt.Println("main thread set up a new task")
	// 启动一个并发任务
    go func(ctx context.Context) {
        pc, _, _, _ := runtime.Caller(1)
        currentFunction := runtime.FuncForPC(pc).Name()
        fmt.Println("[", currentFunction, "] ", "start")

		for {
			select {
			case <-ctx.Done():
                fmt.Println("[", currentFunction, "] ", "Task canceled")
				return
			default:
				// 模拟一些耗时的操作
				time.Sleep(1 * time.Second)
                fmt.Println("[", currentFunction, "] ", "Doing some work...")
			}
		}

    }(ctx)

    // 模拟一段时间后取消上下文
    time.Sleep(3 * time.Second)
    fmt.Println("main thread cancel ctx")
    cancel()

    // 等待一段时间以观察输出
    time.Sleep(2 * time.Second)
    fmt.Println("main thread exit")
}