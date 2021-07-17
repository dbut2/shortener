package main

import (
	"github.com/dbut2/shortener/pkg/cli"
)

func main() {
	cli := cli.New("https://but.la")
	cmd := cli.Shorten()
	err := cmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
