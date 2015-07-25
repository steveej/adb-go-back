package main

import "../lib"

type RestoreCommand struct {
	Infile string `short:"i" long:"infile" description:"File to restire" required:"yes"`
}

func (x *RestoreCommand) Execute(args []string) error {
	return lib.Restore(x.Infile)
}

var restoreCommand RestoreCommand

func init() {
	Parser.AddCommand("restore",
		"Restore Apps from a file",
		"",
		&restoreCommand)
}
