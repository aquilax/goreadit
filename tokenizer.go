package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/neurosnap/sentences"
)

type Tokenizer struct {
	sentences []*sentences.Sentence
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Process(fileName string) error {
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
	return nil
}
