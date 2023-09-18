package dia

import (
	"sync"

	"github.com/kaloseia/diasmos/domain/dia/yaml"
)

var registry *Registry
var once sync.Once

// GetRegistry is used to get or lazy init the singleton's Registry instance.
func GetRegistry() *Registry {
	once.Do(func() {
		registry = &Registry{
			Models:   make(map[string]yaml.Model),
			Entities: make(map[string]yaml.Entity),
		}
	})
	return registry
}

// ResetRegistry is used to reset the registry for testing purposes.
func ResetRegistry() {
	registry = nil
}
