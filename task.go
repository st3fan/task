// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package task

import "sync"

type Func func(task *Task)

type Task struct {
	closed  chan struct{}
	waiting sync.WaitGroup
}

func New(fn Func) *Task {
	task := &Task{
		closed: make(chan struct{}),
	}
	task.waiting.Add(1)
	go func() {
		fn(task)
		task.waiting.Done()
	}()
	return task
}

func (t *Task) Signal() {
	close(t.closed)
}

func (t *Task) Wait() {
	t.waiting.Wait()
}

func (t *Task) SignalAndWait() {
	t.Signal()
	t.Wait()
}

func (t *Task) HasBeenClosed() <-chan struct{} {
	return t.closed
}
