package atlas

import (
	"io/ioutil"
)

func (client *Client) FetchResult(url string) (string, error) {
	opts := make(map[string]string)

	client.mergeGlobalOptions(opts)
	req := client.prepareRequest("FETCH", url, opts)

	//log.Printf("req: %#v", req)
	resp, err := client.call(req)
	err = handleAPIResponse(resp)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//log.Printf("json: %#v\n", m)
	return string(body), nil
}
