package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ===== 使用示例 =====
func main() {
	scheduler := NewScheduler(3)

	// 提交任务
	scheduler.Submit(Task{ID: "A", Func: func() { time.Sleep(1 * time.Second) }})
	scheduler.Submit(Task{ID: "B", Func: func() { time.Sleep(2 * time.Second) }})
	scheduler.Submit(Task{ID: "C", Func: func() { time.Sleep(500 * time.Millisecond) }})

	// 等待执行完成
	time.Sleep(3 * time.Second)

	// 安全停止并获取结果
	results := scheduler.Stop()

	for id, duration := range results {
		fmt.Printf("任务 %s 耗时: %v\n", id, duration)
	}
}

type Task struct {
	ID   string
	Func func()
}

type Result struct {
	TaskID string
	Time   time.Duration
}

type Scheduler struct {
	tasks   chan Task
	results chan Result
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc // 新增取消函数字段
}

func NewScheduler(workerNum int) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Scheduler{
		tasks:   make(chan Task, 100),
		results: make(chan Result, 100),
		ctx:     ctx,
		cancel:  cancel, // 保存取消函数
	}

	for i := 0; i < workerNum; i++ {
		s.wg.Add(1)
		go s.worker()
	}
	return s
}

func (s *Scheduler) Submit(task Task) {
	select {
	case s.tasks <- task: // 正常提交任务
	case <-s.ctx.Done(): // 已关闭时忽略提交
	}
}

func (s *Scheduler) worker() {
	defer s.wg.Done()
	for {
		select {
		case <-s.ctx.Done(): // 优先检查取消信号
			return
		case task, ok := <-s.tasks:
			if !ok { // 通道已关闭
				return
			}
			start := time.Now()
			task.Func()
			elapsed := time.Since(start)

			// 发送结果前再次检查上下文
			select {
			case s.results <- Result{TaskID: task.ID, Time: elapsed}:
			case <-s.ctx.Done():
				return
			}
		}
	}
}

// 修复重点：确保安全关闭通道
func (s *Scheduler) Stop() map[string]time.Duration {
	s.cancel()     // 发送取消信号
	close(s.tasks) // 关闭任务通道（解除worker阻塞）

	// 等待所有worker退出（关键！）
	s.wg.Wait()

	close(s.results) // 安全关闭结果通道（此时无worker发送数据）

	// 收集剩余结果
	results := make(map[string]time.Duration)
	for res := range s.results {
		results[res.TaskID] = res.Time
	}
	return results
}
