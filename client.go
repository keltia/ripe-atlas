package atlas

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"time"
)

type context struct {
	config Config
	client *http.Client
}

var (
	// ctx is out internal context
	ctx *context
)

// NewClient is the first function to call.
// Yes, it does take multiple config
// and the last one wins.
func NewClient(cfgs ...Config) (*Client, error) {
	client := &Client{}
	ctx = &context{}
	for _, cfg := range cfgs {
		ctx.config = cfg
	}
	ctx.client = addHTTPClient(ctx)
	return client, nil
}

func getProxy(req *http.Request) (uri *url.URL, err error) {
	uri, err = http.ProxyFromEnvironment(req)
	if err != nil {
		log.Printf("no proxy in environment")
		uri = &url.URL{}
	} else if uri == nil {
		log.Println("No proxy configured or url excluded")
	}
	return
}

func setupTransport(ctx *context) (*http.Transport, error) {
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
	proxyURL, err := getProxy(req)
	if err != nil {
		if ctx.config.Verbose {
			log.Println("no proxy defined")
		}
	}

	if ctx.config.ProxyAuth != "" {
		req.Header.Set("Proxy-Authorization", ctx.config.ProxyAuth)
	}

	transport := &http.Transport{
		Proxy:              http.ProxyURL(proxyURL),
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		ProxyConnectHeader: req.Header,
	}

	return transport, nil
}

func addHTTPClient(ctx *context) *http.Client {
	transport, err := setupTransport(ctx)
	if err != nil {
		log.Fatalf("unable to create httpclient: %v", err)
	}
	return &http.Client{Transport: transport, Timeout: 20 * time.Second}
}
