package main

import (
	"github.com/nsf/termbox-go"
)

type Termbox struct {
	input  *Input
	width  int
	height int
}

func NewTermbox() *Termbox {
	return &Termbox{
		input: NewInput(),
	}
}

func (t *Termbox) Run() error {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputAlt)
	t.updateSize(termbox.Size())
	t.input.Start()

mainloop:
	for {
		select {
		case ev := <-t.input.eventQ:
			if ev.Key == t.input.endKey {
				break mainloop
			} else if ev.Type == termbox.EventResize {
				t.updateSize(ev.Width, ev.Height)
			} else if ev.Type == termbox.EventError {
				panic(ev.Err.Error())
			}
			t.update(ev)
		default:
			t.update(termbox.Event{Type: termbox.EventNone})
		}
		t.draw()
	}
	return nil
}

func (t *Termbox) updateSize(w, h int) {
	t.width = w
	t.height = h
}

func (t *Termbox) update(ev termbox.Event) {

}

func (t *Termbox) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(3, 4, 'a', termbox.ColorWhite, termbox.ColorDefault)
	termbox.Flush()
}
