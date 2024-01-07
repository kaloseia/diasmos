package plugin

import (
	"errors"
	"fmt"
)

var ErrKeyInvalidFormat = errors.New("invalid format: expected <type>/<name>@<version>")

var ErrIDInvalidField = errors.New("invalid id: invalid field")
var ErrIDInvalidFieldType = fmt.Errorf("%w 'Type' (expected: %v)", ErrIDInvalidField, reKeyType)
var ErrIDInvalidFieldName = fmt.Errorf("%w 'Name' (expected: %v)", ErrIDInvalidField, reKeyName)
var ErrIDInvalidFieldVersion = fmt.Errorf("%w 'Version' (expected: %v)", ErrIDInvalidField, reKeyVersion)

func errWrapID(pluginID ID, toWrap error) error {
	return fmt.Errorf("plugin '%s': %w", pluginID, toWrap)
}
