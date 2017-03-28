// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/naoina/toml"
)

// Config holds our parameters
type Config struct {
	APIKey       string
	DefaultProbe int
	PoolSize     int
	WantAF       string
}

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
		// file must be a tag so add a "."
		str = filepath.Join(os.Getenv("HOME"),
			fmt.Sprintf(".%s", file),
			"config.toml")
	}
	return
}

// LoadConfig reads a file as a TOML document and return the structure
func LoadConfig(file string) (c *Config, err error) {
	c = new(Config)

	// Check for tag
	sFile := checkName(file)
	if sFile == "" {
		return c, fmt.Errorf("Wrong format for %s", file)
	}

	// Check if there is any config file
	if _, err := os.Stat(sFile); err != nil {
		// No config file is no error
		return c, nil
	}

	// Read it
	buf, err := ioutil.ReadFile(sFile)
	if err != nil {
		return c, fmt.Errorf("Can not read %s", sFile)
	}

	err = toml.Unmarshal(buf, &c)
	if err != nil {
		return c, fmt.Errorf("Error parsing toml %s: %v",
			sFile, err)
	}

	return c, nil
}
