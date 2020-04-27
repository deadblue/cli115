package spinner

import (
	"fmt"
	"io"
	"time"
)

type Spinner interface {
	Start()
	Stop()
}

type implSpinner struct {
	alive bool
	ch    chan struct{}

	out      io.Writer
	interval time.Duration
	frames   []string
	prefix   string
	suffix   string
	complete string
}

func (s *implSpinner) Start() {
	s.alive = true
	s.ch = make(chan struct{}, 1)
	go s.run()
}

func (s *implSpinner) run() {
	count := len(s.frames)
	for f := 0; s.alive; {
		frame := s.frames[f%count]
		f += 1
		fmt.Printf("\r%s%s%s", s.prefix, frame, s.suffix)
		time.Sleep(s.interval)
	}
	// clear current line
	fmt.Print("\x1b[2K\n\x1b[1A")
	fmt.Println(s.complete)
	close(s.ch)
}

func (s *implSpinner) Stop() {
	s.alive = false
	// Waiting for complete done
	<-s.ch
}
