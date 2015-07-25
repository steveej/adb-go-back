package main

import (
	"log"

	"../lib"
)

type BackupCommand struct {
	Outfile string `short:"o" long:"outfile" description:"Output filename" required:"yes"`
	Infile  string `short:"i" long:"infile" description:"Infile filename" required:"yes"`
}

func (x *BackupCommand) Execute(args []string) error {
	apps, err := lib.LoadYaml(x.Infile)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return apps.Backup(x.Outfile, "-apk", "-obb", "-noshared", "-nosystem")
}

var backupCommand BackupCommand

func init() {
	Parser.AddCommand("backup",
		"Backup Apps to a file",
		"Backup the Apps listed in Infile to Outfile",
		&backupCommand)
}
