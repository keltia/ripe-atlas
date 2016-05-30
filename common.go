// common.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package atlas

import ()
import "github.com/bndr/gopencils"

const (
	apiEndpoint = "https://atlas.ripe.net/api/v2"
)

var (
	// APIUser is the RIPE username
	APIUser string
	// APIPassword is the RIPE user password
	APIPassword string
)

// SetAuth stores the credentials for later use
func SetAuth(user, pwd string) {
	APIUser = user
	APIPassword = pwd
}

// WantAuth returns either a BasicAuth or nil depending on stored credentials
func WantAuth() (auth *gopencils.BasicAuth) {
	if APIUser == "" || APIPassword == "" {
		return nil
	}
	auth = &gopencils.BasicAuth{
		Username: APIUser,
		Password: APIPassword,
	}
	return auth
}
