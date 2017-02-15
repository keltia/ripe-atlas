package main

import (
    "github.com/urfave/cli"
    "github.com/keltia/ripe-atlas"
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
                Action: keysList,
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

// probeList displays all probes
func keysList(c *cli.Context) (err error) {
    kl, err := atlas.GetKeys()
    if err != nil {
        log.Printf("GetKeys err: %v - kl:%v", err, kl)
        os.Exit(1)
    }
    log.Printf("Got %d keys\n/#v", len(kl), kl)
    return
}

func keysInfo(c *cli.Context) (err error) {
    return
}
