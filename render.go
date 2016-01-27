package main

type Render interface {
	Run(t *Tokenizer) error
}
