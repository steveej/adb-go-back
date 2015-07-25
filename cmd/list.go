package main

import (
	"fmt"
	"log"

	"../lib"
)

type ListCommand struct {
}

var listCommand ListCommand

func (x *ListCommand) Execute(args []string) error {
	apps, err := lib.GetUserApps()
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	fmt.Printf("%v\n", apps)
	return err
}

func init() {
	Parser.AddCommand("list",
		"List all Apps",
		"This command lists the Apps installed on the ADB device.",
		&listCommand)
}
