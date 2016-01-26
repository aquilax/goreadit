package main

import "github.com/nsf/termbox-go"

type Input struct {
	endKey termbox.Key
	eventQ chan termbox.Event
	ctrl   chan bool
}

func NewInput() *Input {
	return &Input{
		endKey: termbox.KeyCtrlC,
		eventQ: make(chan termbox.Event),
		ctrl:   make(chan bool, 2),
	}
}

func (i *Input) Start() {
	go poll(i)
}

func (i *Input) Stop() {
	i.ctrl <- true
}

func poll(i *Input) {
loop:
	for {
		select {
		case <-i.ctrl:
			break loop
		default:
			i.eventQ <- termbox.PollEvent()
		}
	}
}
