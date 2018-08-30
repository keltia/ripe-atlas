// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package main

import (
	"fmt"
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

// Config holds our parameters
type Config struct {
	APIKey       string
	DefaultProbe int
	ProxyAuth    string

	ProbeSet struct {
		PoolSize int
		Type     string
		Value    string
		Tags     string
	}

	Measurements struct {
		BillTo string
	}
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
