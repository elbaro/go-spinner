package spinner

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var chars = []rune("⣾⣽⣻⢿⡿⣟⣯⣷")

type Spinner struct {
	lock       *sync.Mutex
	msg        string
	lastOutLen int
	stopSignal chan struct{}
}

func New(msg string) *Spinner {
	s := &Spinner{
		lock:       &sync.Mutex{},
		msg:        msg,
		lastOutLen: 0,
		stopSignal: make(chan struct{}, 1),
	}

	go func(s *Spinner) {
		for {
			for i := 0; i < 8; i++ {
				select {
				case <-s.stopSignal:
					return
				default:
					line := fmt.Sprintf("%c %s", chars[i], s.msg)
					line_len := utf8.RuneCountInString(line)

					s.lock.Lock()
					s.erase()
					fmt.Print(line)
					s.lastOutLen = line_len
					s.lock.Unlock()
				}
				time.Sleep(125 * time.Millisecond)
			}
		}
	}(s)

	return s
}

func (s *Spinner) erase() {
	fmt.Print(strings.Repeat("\b", s.lastOutLen))
}

func (s *Spinner) Stop() {
	s.stopSignal <- struct{}{}
	s.lock.Lock()
	s.erase()
	fmt.Printf("✔ %s\n", s.msg)
	s.lock.Unlock()
}

func main() {
	s := New("hello ..")
	time.Sleep(3 * time.Second)
	s.Stop()
}
