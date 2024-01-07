package plugin

import (
	"fmt"
	"regexp"
)

type ID struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    Type   `json:"type"`
}

var reKey = fmt.Sprintf("^%s/%s@%s$", reKeyType, reKeyName, reKeyVersion)
var reKeyType = `(compile|execute)`
var reKeyName = `([a-zA-Z0-9\-]+)`
var reKeyVersion = `(\d+\.\d+\.\d+)`

// FromKey parses a string plugin identifier key in the format `<type>/<name>@<version>`.
//
// Example key: compile/go-struct@5.3.2
func (id *ID) FromKey(key string) error {
	pattern := regexp.MustCompile(reKey)
	matches := pattern.FindStringSubmatch(key)

	if matches == nil || len(matches) != 4 {
		return ErrKeyInvalidFormat
	}

	id.Type = Type(matches[1])
	id.Name = matches[2]
	id.Version = matches[3]

	return nil
}

// ToKey validates and outputs a plugin identifier as a plugin identifier key in the format `<type>/<name>@<version>`.
//
// Example key: compile/go-struct@5.3.2
func (id *ID) ToKey() (string, error) {
	validErr := id.Validate()
	if validErr != nil {
		return "", validErr
	}
	return id.String(), nil
}

// String implements the Stringer interface and outputs the plugin identifier as a plugin identifier key in the format `<type>/<name>@<version>` without validation.
func (id ID) String() string {
	return fmt.Sprintf("%v/%s@%s", id.Type, id.Name, id.Version)
}

// Validate checks the ID's fields and values for completeness and having the correct string format.
func (id *ID) Validate() error {
	patternType := regexp.MustCompile("^" + reKeyType + "$")
	if id.Type == "" || !patternType.MatchString(string(id.Type)) {
		return ErrIDInvalidFieldType
	}
	patternName := regexp.MustCompile("^" + reKeyName + "$")
	if id.Name == "" || !patternName.MatchString(string(id.Name)) {
		return ErrIDInvalidFieldName
	}
	patternVersion := regexp.MustCompile("^" + reKeyVersion + "$")
	if id.Version == "" || !patternVersion.MatchString(string(id.Version)) {
		return ErrIDInvalidFieldVersion
	}

	return nil
}
