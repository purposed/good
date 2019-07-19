package timer

import (
	"fmt"
	"strings"
	"time"
)

// StackTimer allows to get the timing breakdown of a nested process.
type StackTimer struct {
	name  string
	stack *timeStack

	currentBuffer strings.Builder
}

// NewStackTimer returns a new stack timer.
func NewStackTimer(name string) *StackTimer {
	return &StackTimer{
		stack: &timeStack{},
	}
}

// Start pushes an entry on the stack.
func (t *StackTimer) Start(name string) {
	if t == nil {
		return
	}

	var paddingStr []string
	for i := 0; i < t.stack.Depth(); i++ {
		paddingStr = append(paddingStr, "|\t")
	}

	if _, err := t.currentBuffer.Write([]byte(
		fmt.Sprintf("%s%s\n", strings.Join(paddingStr, ""), name))); err != nil {
		panic(err)
	}
	t.stack.Push(name, time.Now())
}

// End pops an entry from the stack.
func (t *StackTimer) End() {
	if t == nil {
		return
	}

	en, err := t.stack.Pop()
	if err != nil {
		panic(err)
	}

	var paddingStr []string
	for i := 0; i < t.stack.Depth(); i++ {
		paddingStr = append(paddingStr, "|\t")
	}

	if _, err := t.currentBuffer.Write([]byte(
		fmt.Sprintf("%s~> Took %f ms\n", strings.Join(paddingStr, ""), time.Since(en.time).Seconds()*1000))); err != nil {
		panic(err)
	}
}

// Trace returns the current stack trace.
func (t *StackTimer) Trace() string {
	if t == nil {
		return ""
	}
	return t.currentBuffer.String()
}
