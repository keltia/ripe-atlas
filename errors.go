package atlas

import "errors"

// ErrInvalidMeasurementType is a new error
var ErrInvalidMeasurementType = errors.New("invalid measurement type")

// ErrInvalidAPIKey is returned when the key is invalid
var ErrInvalidAPIKey = errors.New("invalid API key")

// ErrAPIKeyIsMandatory is returned when a call need one
var ErrAPIKeyIsMandatory = errors.New("API call requires an API key")
