// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Check the parameter for either tag or filename
func checkName(file string) (str string) {
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

	// Check for tag
	if !strings.HasSuffix(file, ".toml") {
		str = filepath.Join(basedir, file, "config.toml")
	}
	return
}
