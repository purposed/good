package task

import (
	"sync"
	"time"

	"github.com/purposed/good"
)

// Parameters are used to configure a task.
type Parameters struct {
	Name     string
	Function func() error
	Logger   good.Logger
}

// RecurringTask represents an asychronous repeating task.
type RecurringTask struct {
	Name string
	fn   func() error

	log good.Logger

	trigger chan bool

	stop chan bool
	wg   sync.WaitGroup
}

// New returns a new recurring task.
func New(p Parameters) *RecurringTask {
	logger := p.Logger
	if p.Logger == nil {
		logger = good.DefaultLogger
	}
	return &RecurringTask{
		Name:    p.Name,
		fn:      p.Function,
		stop:    make(chan bool),
		trigger: make(chan bool),
		log:     logger,
	}
}

func (t *RecurringTask) do() {
	if err := t.fn(); err != nil {
		t.log.Errorf("error in task: %s", err.Error())
	} else {
		t.log.Info("task complete")
	}
}

// Start starts the task in the background, running it
// at the defined interval.
func (t *RecurringTask) Start(interval time.Duration) {
	t.wg.Add(1)

	go func() {
		var shouldStop bool

		for !shouldStop {
			select {
			case <-t.trigger:
				t.log.Info("task triggered manually")
				t.do()
			case <-time.After(interval):
				t.log.Info("running task")
				t.do()
			case <-t.stop:
				shouldStop = true
			}
		}
		t.wg.Done()
	}()
}

// Trigger triggers the task manually.
func (t *RecurringTask) Trigger() {
	t.trigger <- true
}

// Stop stops the task routine.
func (t *RecurringTask) Stop() {
	t.stop <- true
	t.wg.Wait()
}
