package atlas

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
	"os"
)

// NewClient is the first function to call.
// Yes, it does take multiple config
// and the last one wins.
func NewClient(cfgs ...Config) (*Client, error) {
	client := &Client{}
	for _, cfg := range cfgs {
		client.config = cfg
	}

	// This holds the global options
	client.opts = make(map[string]string)

	// If no log output is specified, use the default one
	if client.log == nil {
		client.log = log.New(os.Stderr, "ripe-atlas", log.LstdFlags | log.LUTC)
	}

	// Create and save the http.Client
	return client.addHTTPClient()
}

// HasAPIKey returns whether an API key is stored
func (client *Client) HasAPIKey() (string, bool) {
	if client.config.APIKey == "" {
		return "", false
	}
	return client.config.APIKey, true
}

// call is s shortcut
func (client *Client) call(req *http.Request) (*http.Response, error) {
	if client.config.Verbose {
		log.Printf("Full URL:\n%v", req.URL)
	}

	return client.client.Do(req)
}

func (client *Client) setupTransport() (*http.Transport, error) {
	/*
	   Proxy code taken from https://github.com/LeoCBS/poc-proxy-https/blob/master/main.go
	   Analyse endPoint to check proxy stuff
	*/
	req, err := http.NewRequest("HEAD", apiEndpoint, nil)
	if err != nil {
		log.Printf("error: transport: %v", err)
		return nil, err
	}

	// Get proxy URL
	proxyURL, err := http.ProxyFromEnvironment(req)
	if err != nil {
		if client.config.Verbose {
			log.Println("no proxy defined")
		}
	}

	if client.config.ProxyAuth != "" {
		req.Header.Set("Proxy-Authorization", client.config.ProxyAuth)
	}

	transport := &http.Transport{
		Proxy:              http.ProxyURL(proxyURL),
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		ProxyConnectHeader: req.Header,
	}

	return transport, nil
}

func (client *Client) addHTTPClient() (*Client, error) {
	transport, err := client.setupTransport()
	if err != nil {
		log.Fatalf("unable to create httpclient: %v", err)
	}
	client.client = &http.Client{Transport: transport, Timeout: 20 * time.Second}
	return client, err
}

// SetAF sets the global option WantAF
func (client *Client) SetAF(family string) *Client {
	return client.SetOption("wantAF", family)
}

// SetFormat sets the global option
func (client *Client) SetFormat(format string) *Client {
	return client.SetOption("format", format)
}

// SetInclude sets the global option
func (client *Client) SetInclude(include string) *Client {
	return client.SetOption("include", include)
}

// SetOption sets a global option
func (client *Client) SetOption(name, value string) *Client {
	if value != "" {
		client.opts[name] = value
	}
	return client
}

func (client *Client) mergeGlobalOptions(opts map[string]string) {
	for k, v := range client.opts {
		opts[k] = v
	}
}
