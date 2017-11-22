// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"errors"
	"github.com/naoina/toml"
)

const (
	proxyTag = "proxy"
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

func setupProxyAuth() (auth string, err error) {
	if fDebug {
		log.Printf("Looking for proxy credentials in %s:", dbrcFile)
	}

	user, password, err := loadDbrc(dbrcFile)
	if err != nil {
		if fVerbose {
			log.Printf("No dbrc file: %v", err)
		}
	} else {
		// Do we have a proxy user/password?
		if user != "" && password != "" {
			auth = fmt.Sprintf("%s:%s", user, password)
			auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

			if fVerbose {
				log.Printf("Proxy user %s found.", user)
			}
		} else {
			auth = "nothing"
			err = errors.New("invalid proxy creds")
		}
	}
	return
}

func loadDbrc(file string) (user, password string, err error) {
	fh, err := os.Open(file)
	if err != nil {
		return "", "", fmt.Errorf("error: can not find %s: %v", file, err)
	}
	defer fh.Close()

	/*
	   Format:
	   <db>     <user>    <pass>   <type>
	*/
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		// Replace all tabs by a single space
		l := strings.Replace(line, "\t", " ", -1)
		flds := strings.Split(l, " ")

		// Check what we need
		if flds[0] == proxyTag {
			user = flds[1]
			password = flds[2]
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("reading dbrc %s", dbrcFile)
	}

	if user == "" {
		return "", "", fmt.Errorf("no user/password for %s in %s", proxyTag, dbrcFile)
	}

	return
}
