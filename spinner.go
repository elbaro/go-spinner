package spinner

import (
	"fmt"
	"sync"
	"time"
)

var chars = []rune("⣾⣽⣻⢿⡿⣟⣯⣷")
var eraser = "\x1B[2K\r"

type Spinner struct {
	lock       *sync.Mutex
	msg        string
	stopSignal chan struct{}
}

func New(msg string) *Spinner {
	s := &Spinner{
		lock:       &sync.Mutex{},
		msg:        msg,
		stopSignal: make(chan struct{}, 1),
	}

	go func(s *Spinner) {
		for {
			for i := 0; i < 8; i++ {
				s.lock.Lock()
				select {
				case <-s.stopSignal:
					s.lock.Unlock()
					return
				default:
					fmt.Printf("%s%c %s", eraser, chars[i], s.msg)
				}
				s.lock.Unlock()
				time.Sleep(125 * time.Millisecond)
			}
		}
	}(s)

	return s
}

func (s *Spinner) Done() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	fmt.Printf("%s✔ %s\n", eraser, s.msg)
	s.lock.Unlock()
}

func (s *Spinner) DoneClean() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	fmt.Printf("%s", eraser)
	s.lock.Unlock()
}

func (s *Spinner) Fail() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	fmt.Printf("%s❌ %s\n", eraser, s.msg)
	s.lock.Unlock()
}

func (s *Spinner) FailClean() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	fmt.Printf("%s", eraser)
	s.lock.Unlock()
}
