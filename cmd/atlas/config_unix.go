// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

// +build unix !windows

package main

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// Default location is now $HOME/.config/<tag>/ on UNIX
	basedir = filepath.Join(os.Getenv("HOME"), ".config", MyName)
)

// Check the parameter for either tag or filename
func checkName(file string) (str string) {
	// Use default location
	if file == "" {
		str = filepath.Join(basedir, file, "config.toml")
	} else {
		// Full path, MUST have .toml
		if bfile := []byte(file); bfile[0] == '/' {
			if !strings.HasSuffix(file, ".toml") {
				str = ""
			} else {
				str = file
			}
			return
		}

		// If ending with .toml, take it literally
		if strings.HasSuffix(file, ".toml") {
			str = file
			return
		}
	}
	return
}
