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

		filePathAbs := filepath.Join(dirPath, fileName)
		fileLoadErr := registry.loadModelFromFile(filePathAbs)
		if fileLoadErr != nil {
			return fileLoadErr
		}
	}
	return nil
}

func (registry *Registry) loadModelFromFile(filePathAbs string) error {
	var modelDefinition yaml.Model
	unmarshalErr := unmarshalYAMLFile(&modelDefinition, filePathAbs)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	_, nameConflict := registry.Models[modelDefinition.Name]
	if nameConflict {
		return fmt.Errorf("model name '%s' already exists in registry (conflict: %s)", modelDefinition.Name, filePathAbs)
	}

	registry.Models[modelDefinition.Name] = modelDefinition
	return nil
}

func unmarshalYAMLFile[TTarget any](target *TTarget, filePathAbs string) error {
	fileContents, readFileErr := os.ReadFile(filePathAbs)
	if readFileErr != nil {
		return fmt.Errorf("error reading file contents '%s': %w", filePathAbs, readFileErr)
	}

	unmarshalErr := yaml3.Unmarshal(fileContents, target)
	if unmarshalErr != nil {
		return fmt.Errorf("error unmarshalling yaml file contents '%s': %w", filePathAbs, unmarshalErr)
	}

	return nil
}
