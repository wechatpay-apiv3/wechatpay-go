// Copyright 2021 Tencent Inc. All rights reserved.

package task

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRepeatedTask(t *testing.T) {
	cnt := new(int)

	task := NewRepeatedTask(1*time.Second, func(time.Time) { *cnt++ })
	assert.Equal(t, Init, task.State())

	task.Start()
	assert.Equal(t, Running, task.State())

	time.Sleep(5 * time.Second)

	task.Stop()
	assert.Equal(t, Stopped, task.State())

	assert.Equal(t, 5, *cnt)

	time.Sleep(2 * time.Second)

	assert.Equal(t, 5, *cnt)
}

func runTask(cnt *int) {
	task := NewRepeatedTask(1*time.Second, func(time.Time) { *cnt++ })

	task.Start()

	time.Sleep(5 * time.Second)
}

func TestRecycleRepeatedTask(t *testing.T) {
	cnt := new(int)

	runTask(cnt)
	runtime.GC()

	assert.Equal(t, 5, *cnt)

	time.Sleep(2 * time.Second)

	assert.Equal(t, 5, *cnt)
}
