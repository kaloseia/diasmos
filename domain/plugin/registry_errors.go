package plugin

import "errors"

var ErrNotInitialized = errors.New("registry has not been initialized")
var ErrAlreadyInitialized = errors.New("registry has already been initialized")
var ErrNoPluginsDirPath = errors.New("registry plugin directory path has not been set")
