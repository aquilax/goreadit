package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

type appState int8

const (
	wpmStep = 10

	stateRunning appState = iota
	statePause
)

type Termbox struct {
	input       *Input
	currentWord string
	width       int
	height      int
	centerX     int
	centerY     int
	wpm         int
	state       appState
}

func NewTermbox() *Termbox {
	return &Termbox{
		input: NewInput(),
		wpm:   400,
		state: statePause,
	}
}

func (t *Termbox) Run(tokenizer *Tokenizer) error {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.OutputNormal)
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
			t.update(ev, tokenizer)
		default:
			t.update(termbox.Event{Type: termbox.EventNone}, tokenizer)
		}
		t.draw()
		time.Sleep(time.Minute / time.Duration(t.wpm))
	}
	return nil
}

func (t *Termbox) updateSize(w, h int) {
	t.width = w
	t.height = h
	t.centerX = t.width / 2
	t.centerY = t.height / 2
}

func (t *Termbox) update(ev termbox.Event, tokenizer *Tokenizer) {
	if ev.Type == termbox.EventKey {
		if ev.Key == termbox.KeyArrowLeft {
			t.wpm -= wpmStep
		}
		if ev.Key == termbox.KeyArrowRight {
			t.wpm += wpmStep
		}
		if ev.Key == termbox.KeySpace {
			if t.state == statePause {
				t.state = stateRunning
			} else {
				t.state = statePause
			}
		}
	}
	if t.state == statePause {
		return
	}
	if ev.Type == termbox.EventNone {
		if word, ok := tokenizer.getNextWord(); ok {
			t.currentWord = word
		}
	}
}

func (t *Termbox) draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	t.renderWireframe()
	t.renderWord(t.currentWord)
	t.renderWpm(t.wpm)
	termbox.Flush()
}

func (t *Termbox) renderWireframe() {
	for y := t.centerY - 2; y <= t.centerY+2; y++ {
		for x := 0; x < t.width; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorWhite)
		}
	}
	termbox.SetCell(t.centerX, t.centerY-1, '|', termbox.ColorRed, termbox.ColorWhite)
	termbox.SetCell(t.centerX, t.centerY+1, '|', termbox.ColorRed, termbox.ColorWhite)
}

func (t *Termbox) renderWord(word string) {
	breakingPoint := t.getBreakingPoint(word)
	i := 0
	for _, runeValue := range word {
		x := t.centerX - breakingPoint + i
		color := termbox.ColorBlack
		if x == t.centerX {
			color = termbox.ColorRed
		}
		termbox.SetCell(x, t.centerY, runeValue, color, termbox.ColorWhite)
		i++
	}
}

func (t *Termbox) renderWpm(wpm int) {
	str := fmt.Sprintf("WPM: %d, Ctrl-C to exit, space to pause", wpm)
	startX := 1
	for i, runeValue := range str {
		termbox.SetCell(startX+i, t.height-1, runeValue, termbox.ColorWhite, termbox.ColorDefault)
	}
}

func (t *Termbox) getBreakingPoint(word string) int {
	word = strings.ToLower(word)
	vowels := "aeiouyаъоуеиюя"
	i := 0
	for _, runeValue := range word {
		if i == 0 {
			i++
			continue
		}
		if strings.IndexRune(vowels, runeValue) != -1 {
			return i
		}
		i++
	}
	return utf8.RuneCountInString(word) / 2
}
