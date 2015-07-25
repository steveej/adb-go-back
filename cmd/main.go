package main

import (
	"os"

	goflags "github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`
}

var options Options
var Parser = goflags.NewParser(&options, goflags.Default)

func main() {
	if _, err := Parser.Parse(); err != nil {
		os.Exit(1)
	}
}
