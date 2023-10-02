// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package internal

import (
	"sync"
	"time"
)

func NewSingleTask(interval time.Duration, do func()) *SingleTask {
	return &SingleTask{
		interval: interval,
		lastRun:  time.Now(),
		do:       do,
	}
}

type SingleTask struct {
	interval time.Duration
	lastRun  time.Time
	do       func()

	scheduled *time.Timer
	working   bool

	mu sync.Mutex
}

func (r *SingleTask) Run() {
	r.runTask(false)
}

func (r *SingleTask) runTask(scheduled bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if scheduled {
		// this is the scheduled task, clear itself
		r.scheduled = nil
	}

	if !r.working {
		// can run task, but rate limit to 1 task/interval
		delta := time.Since(r.lastRun) - r.interval
		if delta < -10*time.Millisecond {
			// task requesting too fast, schedule

			if r.scheduled == nil { // no task scheduled
				r.scheduled = time.AfterFunc(-delta, func() {
					r.runTask(true)
				})
			}
			return
		}

		r.working = true
		r.lastRun = time.Now()
		r.mu.Unlock()

		r.do()

		r.mu.Lock()
		r.working = false
		return
	}

	// has task running, schedule to wait until no task running

	if r.scheduled == nil { // no scheduled task, create one
		r.scheduled = time.AfterFunc(r.interval, func() {
			r.runTask(true)
		})
		return
	}

	// reaching here because:
	//
	//  a) scheduled task has fired, but it's waiting for the lock.
	//  b) scheduled task not fired.
	//
	// in either case, discard this task, let the scheduled task run.
	return
}
