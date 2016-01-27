package main

type ReadIt struct {
	config *Config
}

func NewReadIt(config *Config) *ReadIt {
	return &ReadIt{config}
}

func (r *ReadIt) Run() error {
	tokenizer := NewTokenizer()
	if err := tokenizer.Process(r.config.FileName); err != nil {
		return err
	}
	return NewTermbox().Run(tokenizer)
}
