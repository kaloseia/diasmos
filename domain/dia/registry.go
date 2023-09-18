package dia

import (
	"fmt"
	"os"
	"path/filepath"

	yaml3 "gopkg.in/yaml.v3"

	"github.com/kaloseia/diasmos/domain/dia/yaml"
)

const ModelFileSuffix = ".mod"
const EntityFileSuffix = ".ent"

type Registry struct {
	Models   map[string]yaml.Model  `yaml:"models"`
	Entities map[string]yaml.Entity `yaml:"entities"`
}

func (registry *Registry) LoadModelsFromDirectory(dirPath string) error {
	if registry.Models == nil {
		registry.Models = make(map[string]yaml.Model)
	}

	dirEntries, readErr := os.ReadDir(dirPath)
	if readErr != nil {
		return fmt.Errorf("error reading directory '%s': %w", dirPath, readErr)
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		fileName := dirEntry.Name()
		if filepath.Ext(fileName) != ModelFileSuffix {
			continue
		}

		filePath := filepath.Join(dirPath, fileName)
		fileContents, readFileErr := os.ReadFile(filePath)
		if readFileErr != nil {
			return fmt.Errorf("error reading file contents '%s': %w", filePath, readFileErr)
		}

		var modelDefinition yaml.Model
		unmarshalErr := yaml3.Unmarshal(fileContents, &modelDefinition)
		if unmarshalErr != nil {
			return fmt.Errorf("error unmarshalling yaml file contents '%s': %w", filePath, unmarshalErr)
		}

		_, nameConflict := registry.Models[modelDefinition.Name]
		if nameConflict {
			return fmt.Errorf("model name '%s' already exists in registry (conflict: %s)", modelDefinition.Name, filePath)
		}

		registry.Models[modelDefinition.Name] = modelDefinition

	}
	return nil
}
