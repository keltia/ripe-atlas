package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
)

// init injects our key-related commands
func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name: "keys",
		Aliases: []string{
			"k",
			"key",
		},
		Usage:       "key-related keywords",
		Description: "All the commands for keys",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Aliases:     []string{"ls"},
				Usage:       "lists all keys",
				Description: "displays all keys",
				Action:      keysList,
			},
			{
				Name:        "info",
				Usage:       "info for one key",
				Description: "gives info for one key",
				Action:      keysInfo,
			},
		},
	})
}

// displayKey display short or verbose data about a key
func displayKey(key *atlas.Key, verbose bool) (res string) {
	if verbose {
		res = fmt.Sprintf("%v\n", key)
	} else {
		res = fmt.Sprintf("UUID: %s Type: %s Active? %v Created: %s\n",
			key.Label,
			key.Type,
			key.IsActive,
			key.CreatedAt)
	}
	return
}

// displayAllKeys display short or verbose data about keys
func displayAllKeys(keys *[]atlas.Key, verbose bool) (res string) {
	res = ""
	for _, key := range *keys {
		res += displayKey(&key, verbose)
	}
	return
}

// probeList displays all probes
func keysList(c *cli.Context) (err error) {
	opts := make(map[string]string)

	// Check global parameters
	opts = checkGlobalFlags(opts)

	if fVerbose {
		displayOptions(opts)
	}

	kl, err := atlas.GetKeys(opts)
	if err != nil {
		log.Printf("GetKeys err: %v - kl:%v", err, kl)
		os.Exit(1)
	}
	log.Printf("Got %d keys\n", len(kl))
	fmt.Print(displayAllKeys(&kl, fVerbose))
	return
}

func keysInfo(c *cli.Context) (err error) {
	args := c.Args()
	if args[0] == "" {
		log.Fatalf("Error: you must specify a UUID!")
	}

	k, err := atlas.GetKey(args[0])
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Print(displayKey(&k, fVerbose))
	return
}
