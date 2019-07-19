package task_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/purposed/good/task"
)

type mockTask struct {
	CallCount  int
	ShouldFail bool

	rw sync.RWMutex
}

func (m *mockTask) Run() error {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.CallCount++
	if m.ShouldFail {
		return errors.New("oops")
	}
	return nil
}

func (m *mockTask) GetCallCount() int {
	m.rw.RLock()
	defer m.rw.RUnlock()
	return m.CallCount
}

func Test_NewTask(t *testing.T) {
	taskDef := &mockTask{}

	task := task.New(task.Parameters{Name: "mock", Function: taskDef.Run})

	if task == nil {
		t.Error("NewTask() returned nil task")
	}
}

func TestTask_Start_Stop(t *testing.T) {
	taskDef := &mockTask{}

	task := task.New(task.Parameters{Name: "mock", Function: taskDef.Run})

	// Schedules task to run every two seconds
	task.Start(200 * time.Millisecond)

	time.Sleep(250 * time.Millisecond)

	if taskDef.GetCallCount() != 1 {
		t.Errorf("task wasn't called")
		return
	}

	time.Sleep(200 * time.Millisecond)

	if taskDef.GetCallCount() != 2 {
		t.Errorf("task wasn't called recurring")
		return
	}

	task.Stop()

	time.Sleep(250 * time.Millisecond)

	if taskDef.GetCallCount() > 2 {
		t.Errorf("task did not stop when asked")
		return
	}
}

func TestTask_Start_Errors(t *testing.T) {
	// Validate that errors don't affect scheduling.

	taskDef := &mockTask{ShouldFail: true}

	task := task.New(task.Parameters{Name: "mock", Function: taskDef.Run})

	// Schedules task to run every two seconds
	task.Start(200 * time.Millisecond)

	time.Sleep(250 * time.Millisecond)

	if taskDef.GetCallCount() != 1 {
		t.Errorf("task wasn't called")
		return
	}

	time.Sleep(200 * time.Millisecond)

	if taskDef.GetCallCount() != 2 {
		t.Errorf("task wasn't called recurring")
		return
	}

	task.Stop()

	time.Sleep(250 * time.Millisecond)

	if taskDef.GetCallCount() > 2 {
		t.Errorf("task did not stop when asked")
		return
	}
}
