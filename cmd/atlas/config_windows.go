// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Check the parameter for either tag or filename
func checkName(file string) (str string) {
	// Full path, begin by \\ or N: MUST have .toml
	bfile := []byte(file)
	if bfile[1] == ':' || strings.HasPrefix(file, "\\\\"){
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

	// Check for tag
	if !strings.HasSuffix(file, ".toml") {
		// file must be a tag so add a "."
		str = filepath.Join(os.Getenv("%LOCALAPPDATA%"),
			"RIPE-Atlas",
			"config.toml")
	}
	return
}
