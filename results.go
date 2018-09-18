package atlas

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// FetchResult downloads result for a given measurement
func (c *Client) FetchResult(url string) (string, error) {
	opts := make(map[string]string)

	c.mergeGlobalOptions(opts)

	// Remove our key for fetching the results
	delete(opts, "key")

	req := c.prepareRequest("FETCH", url, opts)

	c.debug("req=%#v", req)
	c.debug("url=%s", req.URL.String())

	resp, err := c.call(req)
	err = c.handleAPIResponse(resp)
	c.debug("resp=%#v", resp)
	if err != nil {
		return "", errors.Wrap(err, "fetchresult/call")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(body), nil
}
