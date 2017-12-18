package atlas

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// NewClient is the first function to call.
// Yes, it does take multiple config
// and the last one wins.
func NewClient(cfgs ...Config) (*Client, error) {
	c := &Client{}
	for _, cfg := range cfgs {
		c.config = cfg
	}

	// This holds the global options
	c.opts = make(map[string]string)

	// If no log output is specified, use the default one
	if c.config.Log == nil {
		c.log = log.New(os.Stderr, "", log.LstdFlags|log.LUTC)
	} else {
		c.log = c.config.Log
	}

	// Create and save the http.Client
	return c.addHTTPClient()
}

// HasAPIKey returns whether an API key is stored
func (c *Client) HasAPIKey() (string, bool) {
	if c.config.APIKey == "" {
		return "", false
	}
	return c.config.APIKey, true
}

// call is s shortcut
func (c *Client) call(req *http.Request) (*http.Response, error) {
	c.verbose("Full URL:\n%v", req.URL)

	return c.client.Do(req)
}

func (c *Client) setupTransport() (*http.Transport, error) {
	/*
	   Proxy code taken from https://github.com/LeoCBS/poc-proxy-https/blob/master/main.go
	   Analyse endPoint to check proxy stuff
	*/
	req, err := http.NewRequest("HEAD", apiEndpoint, nil)
	if err != nil {
		c.log.Printf("error: transport: %v", err)
		return nil, err
	}

	// Get proxy URL
	proxyURL, err := http.ProxyFromEnvironment(req)
	if err != nil {
		c.verbose("no proxy defined")
	}

	if c.config.ProxyAuth != "" {
		req.Header.Set("Proxy-Authorization", c.config.ProxyAuth)
	}

	myurl, _ := url.Parse(apiEndpoint)
	req.Header.Set("Host", myurl.Host)
	req.Header.Set("User-Agent", fmt.Sprintf("ripe-atlas/%s", ourVersion))

	transport := &http.Transport{
		Proxy:              http.ProxyURL(proxyURL),
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		ProxyConnectHeader: req.Header,
	}

	return transport, nil
}

func (c *Client) addHTTPClient() (*Client, error) {
	transport, err := c.setupTransport()
	if err != nil {
		c.log.Fatalf("unable to create httpclient: %v", err)
	}
	c.client = &http.Client{Transport: transport, Timeout: 20 * time.Second}
	return c, err
}

// SetOption sets a global option
func (c *Client) SetOption(name, value string) *Client {
	if value != "" {
		c.opts[name] = value
	}
	return c
}

func (c *Client) mergeGlobalOptions(opts map[string]string) {
	for k, v := range c.opts {
		opts[k] = v
	}
}
