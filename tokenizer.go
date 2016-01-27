package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/neurosnap/sentences"
)

type Tokenizer struct {
	sentences       []*sentences.Sentence
	currentSentence []string
	sentenceId      int
	wordId          int
	finished        bool
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Process(fileName string) error {
	t.sentenceId = 0
	t.wordId = 0
	t.finished = false
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return t.process(file)
}

func (t *Tokenizer) process(reader io.Reader) error {
	// TODO: Use proper training data
	training := sentences.NewStorage()
	tokenizer := sentences.NewSentenceTokenizer(training)
	text, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	t.sentences = tokenizer.Tokenize(string(text))
	if len(t.sentences) > 0 {
		t.currentSentence = t.processSentence(t.sentences[t.sentenceId].String())
	}
	return nil
}

func (t *Tokenizer) processSentence(text string) []string {
	return strings.Fields(text)
}

func (t *Tokenizer) getNextWord() (string, bool) {
	if t.finished {
		return "", false
	}
	t.wordId++
	if len(t.currentSentence) == t.wordId {
		t.sentenceId++
		if len(t.sentences) == t.sentenceId {
			return "", false
		}
		sentence := strings.TrimSpace(t.sentences[t.sentenceId].Text)
		if sentence == "" {
			t.finished = true
			return "", false
		}
		t.currentSentence = t.processSentence(t.sentences[t.sentenceId].Text)
		t.wordId = 0
	}
	return t.currentSentence[t.wordId], true
}
