package atlas

import (
	"io/ioutil"
)

// FetchResult downloads result for a given measurement
func (c *Client) FetchResult(url string) (string, error) {
	opts := make(map[string]string)

	c.mergeGlobalOptions(opts)

	// Remove our key for fetching the results
	delete(opts, "key")

	req := c.prepareRequest("FETCH", url, opts)

	//log.Printf("req: %#v", req)
	resp, err := c.call(req)
	err = c.handleAPIResponsese(resp)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//log.Printf("json: %#v\n", m)
	return string(body), nil
}
