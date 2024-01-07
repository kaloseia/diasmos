package dia

import "github.com/kaloseia/diasmos/domain/dia/yaml"

func NewRegistry() *Registry {
	return &Registry{
		Models:   map[string]yaml.Model{},
		Entities: map[string]yaml.Entity{},
	}
}
