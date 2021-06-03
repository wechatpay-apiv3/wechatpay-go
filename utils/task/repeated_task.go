package task

import (
	"runtime"
	"sync"
	"time"
)

type State int

const (
	Init State = iota
	Running
	Stopped
)

// implRepeatedTask 定时重复任务
// 将会在后台（另一个goroutine中）定时重复执行给定的任务
//
// 注意：
//  1. 请不要在多个协程中操作同一个 implRepeatedTask 实例，它并不支持多协程并发
//  2. 基于上一条，请不要在 handler 中操作本 implRepeatedTask 实例
type implRepeatedTask struct {
	interval time.Duration
	handler  func(time.Time)
	state    State
	ticker   *time.Ticker

	closed chan struct{}
	wg     sync.WaitGroup
}

// State 查看当前任务状态
func (t *implRepeatedTask) State() State {
	return t.state
}

// Interval 查看当前任务间隔
func (t *implRepeatedTask) Interval() time.Duration {
	return t.interval
}

// Start 启动重复任务
// 当且仅当任务并未启动时有效，其他状态下不会发生任何事
func (t *implRepeatedTask) Start() {
	if t.state != Init {
		return
	}

	t.ticker = time.NewTicker(t.interval)
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		t.run()
	}()

	t.state = Running
}

// Stop 停止重复任务
// 当且仅当任务处于 Running 状态时有效，其他状态下不会发生任何事
func (t *implRepeatedTask) Stop() {
	if t.state != Running {
		return
	}

	close(t.closed)
	t.wg.Wait()
	t.ticker.Stop()
	t.state = Stopped
}

func (t *implRepeatedTask) run() {
	for {
		select {
		case <-t.closed:
			return
		case tickTime := <-t.ticker.C:
			t.handler(tickTime)
		}
	}
}

type wrapper struct {
	*implRepeatedTask
}

type RepeatedTask wrapper

func NewRepeatedTask(interval time.Duration, handler func(time.Time)) *RepeatedTask {
	task := implRepeatedTask{
		closed:   make(chan struct{}),
		interval: interval,
		handler:  handler,
	}

	result := &RepeatedTask{&task}

	runtime.SetFinalizer(result, func(t *RepeatedTask) { t.Stop() })

	return result
}
