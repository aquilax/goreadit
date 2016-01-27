package main

type Debug struct{}

func NewDebug() *Debug {
	return &Debug{}
}

func (d *Debug) Run(tokenizer *Tokenizer) error {
	for {
		word, ok := tokenizer.getNextWord()
		if !ok {
			return nil
		}
		println(word)
	}
	return nil
}
