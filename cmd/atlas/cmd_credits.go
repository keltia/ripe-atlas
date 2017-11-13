package main

import (
	"github.com/urfave/cli"
	"os"
	"log"
	"fmt"
)

// init injects our key-related commands
func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name: "credits",
		Aliases: []string{
			"c",
		},
		Usage:       "credits-related keywords",
		Description: "All the commands for credits",
		Action:      creditsList,
	})
}

func creditsList(c *cli.Context) {
	cl, err := client.GetCredits()
	if err != nil {
		log.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Credits:\n%v", cl)
}
