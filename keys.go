// keys.go

// This file implements the keys API calls

package atlas

import (
)
import (
    "encoding/json"
    "fmt"
    "github.com/sendgrid/rest"
)

// GetKey returns a given API key
func GetKey(uuid string) (k Key, err error) {
    keyEP := apiEndpoint + "/keys"

    key, ok := HasAPIKey()

    // Add at least one option, the APIkey if present
    hdrs := make(map[string]string)
    opts := make(map[string]string)

    if ok {
        opts["key"] = key
    }

    req := rest.Request{
        BaseURL:     keyEP + fmt.Sprintf("/%d", uuid),
        Method:      rest.Get,
        Headers:     hdrs,
        QueryParams: opts,
    }

    //log.Printf("req: %#v", req)
    r, err := rest.API(req)
    if err != nil {
        err = fmt.Errorf("err: %v - r:%v\n", err, r)
        return
    }

    k = Key{}
    err = json.Unmarshal([]byte(r.Body), &k)
    //log.Printf("json: %#v\n", p)
    return
}

// GetKeys returns all your API keys
func GetKeys() (kl []Key, err error) {
    keyEP := apiEndpoint + "/keys"

    key, ok := HasAPIKey()

    // Add at least one option, the APIkey if present
    hdrs := make(map[string]string)
    opts := make(map[string]string)

    if ok {
        opts["key"] = key
    }

    req := rest.Request{
        BaseURL:     keyEP,
        Method:      rest.Get,
        Headers:     hdrs,
        QueryParams: opts,
    }

    //log.Printf("req: %#v", req)
    r, err := rest.API(req)
    if err != nil {
        err = fmt.Errorf("err: %v - r:%v\n", err, r)
        return
    }

    kl = []Key{}
    fmt.Printf("r: %#v", r.Body)
    err = json.Unmarshal([]byte(r.Body), &kl)
    //log.Printf("json: %#v\n", p)
    return
}
