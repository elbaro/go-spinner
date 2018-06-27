package spinner

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ProgressSpinner struct {
	lock       *sync.Mutex
	msg        string
	stopSignal chan struct{}

	status  *atomic.Value
	display string

	done  int32
	fail  int32
	total int32
}

func NewProgress(msg string) *ProgressSpinner {
	s := &ProgressSpinner{
		lock:       &sync.Mutex{},
		msg:        msg,
		stopSignal: make(chan struct{}, 1),

		status:  &atomic.Value{},
		display: "",

		done:  0,
		fail:  0,
		total: 0,
	}
	s.status.Store("")
	s.updateDisplay()

	go func(s *ProgressSpinner) {
		for {
			for i := 0; i < 8; i++ {
				s.lock.Lock()
				select {
				case <-s.stopSignal:
					s.lock.Unlock()
					return
				default:
					fmt.Printf("%s%c %s", eraser, chars[i], s.display)
				}
				s.lock.Unlock()
				time.Sleep(125 * time.Millisecond)
			}
		}
	}(s)

	return s
}

func (s *ProgressSpinner) updateDisplay() {
	done := atomic.LoadInt32(&s.done)
	fail := atomic.LoadInt32(&s.fail)
	total := atomic.LoadInt32(&s.total)
	s.display = fmt.Sprintf("%s [%d/%d, fail %d] %s", s.msg, done, total, fail, s.status.Load())
}

func (s *ProgressSpinner) SetStatus(status string) {
	s.status.Store(status)
	s.updateDisplay()
}

func (s *ProgressSpinner) Add() {
	atomic.AddInt32(&s.total, 1)
	s.updateDisplay()
}

func (s *ProgressSpinner) Done(msg string) {
	atomic.AddInt32(&s.done, 1)
	s.updateDisplay()
	s.lock.Lock()
	fmt.Printf("%s✔ %s\n", eraser, msg)
	s.lock.Unlock()
}

func (s *ProgressSpinner) DoneClean() {
	atomic.AddInt32(&s.done, 1)
	s.updateDisplay()
}

func (s *ProgressSpinner) Fail(msg string) {
	atomic.AddInt32(&s.fail, 1)
	s.updateDisplay()
	s.lock.Lock()
	fmt.Printf("%s❌ %s\n", eraser, msg)
	s.lock.Unlock()
}

func (s *ProgressSpinner) FailClean() {
	atomic.AddInt32(&s.fail, 1)
	s.updateDisplay()
}

func (s *ProgressSpinner) Stop() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	fmt.Printf("%s✔✔ %s\n", eraser, s.msg)
	s.lock.Unlock()
}

func (s *ProgressSpinner) StopClean() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	fmt.Printf("%s", eraser)
	s.lock.Unlock()
}
