package main

type ReadIt struct {
	config *Config
}

func NewReadIt(config *Config) *ReadIt {
	return &ReadIt{config}
}

func (r *ReadIt) Run() error {
	return nil
}
