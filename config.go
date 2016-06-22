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
	DefProbes int
	ApiKey string
}

// Check the parameter for either tag or filename
func checkName(file string) string {
	// Full path, MUST have .toml
	if bfile := []byte(file); bfile[0] == '/' {
		if !strings.HasSuffix(file, ".toml") {
			return ""
		}
		return file
	}
	// Check for tag
	if !strings.HasSuffix(file, ".toml") {
		// file must be a tag so add a "."
		return filepath.Join(os.Getenv("HOME"),
			fmt.Sprintf(".%s", file),
			"config.toml")
	}
	return file
}

// LoadConfig reads a file as a TOML document and return the structure
func LoadConfig(file string) (c *Config, err error) {
	// Check for tag
	sFile := checkName(file)

	c = new(Config)
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
