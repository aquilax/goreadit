package main

type Config struct {
	FileName string
}

func NewConfig(fileName string) *Config {
	return &Config{
		FileName: fileName,
	}
}
