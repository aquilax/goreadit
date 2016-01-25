package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const (
	appName    = "GoReadIt"
	appUsage   = "Command line speed reading tool"
	appVersion = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion

	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
			fmt.Println("Error: Please provide file name")
			os.Exit(1)
		}
		fileName := c.Args()[0]
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Printf("Error: File %s not found\n", fileName)
			os.Exit(1)
		}
		config := NewConfig(fileName)
		if err := NewReadIt(config).Run(); err != nil {
			panic(err)
		}
	}
	app.Run(os.Args)
}
