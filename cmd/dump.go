package main

import (
	"log"

	"../lib"
)

type DumpCommand struct {
	Outfile string `short:"o" long:"outfile" description:"Output filename" required:"yes"`
}

func (x *DumpCommand) Execute(args []string) error {
	apps, err := lib.GetUserApps()
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return apps.DumpYaml(x.Outfile)
}

var dumpCommand DumpCommand

func init() {
	Parser.AddCommand("dump",
		"Dump Packagelist to file",
		"Dump the Packagelist to a file in YAML format",
		&dumpCommand)
}
