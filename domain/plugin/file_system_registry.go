package plugin

import (
	"encoding/json"
	"os"
	"path"
)

// FileSystemRegistry implements plugin.Registry using a local directory of plugins.
type FileSystemRegistry struct {
	PluginsDirPath string
	isInitialized  bool
}

func (r *FileSystemRegistry) Initialize() error {
	if r.PluginsDirPath == "" {
		return ErrNoPluginsDirPath
	}
	if r.isInitialized {
		return ErrAlreadyInitialized
	}
	r.isInitialized = true
	return nil
}

func (r *FileSystemRegistry) IsInitialized() bool {
	return r.isInitialized
}

// ValidatePlugin validates that the prerequisites of a plugin in the registry are given for the plugin's manifest.Prerequisites may be defined as dependencies, versions, and licensing.
//
// Not to be confused with validating the prerequisites of a plugin in the project during plugin installation.
func (r *FileSystemRegistry) ValidatePlugin(pluginID ID) (Manifest, error) {
	if !r.IsInitialized() {
		return Manifest{}, errWrapID(pluginID, ErrNotInitialized)
	}

	manifest, manifestErr := r.getPluginManifest(pluginID)
	if manifestErr != nil {
		return Manifest{}, errWrapID(pluginID, manifestErr)
	}

	return manifest, nil
}

// LoadPlugin loads a plugin from the registry into the local environment and memory via the passed plugin ID.
func (r *FileSystemRegistry) LoadPlugin(manifest Manifest) (Plugin, error) {
	return Plugin{}, nil
}

func (r *FileSystemRegistry) getPluginManifest(pluginID ID) (Manifest, error) {
	manifestPath, pathErr := r.getPluginManifestPath(pluginID)
	if pathErr != nil {
		return Manifest{}, pathErr
	}
	manifestContents, readErr := os.ReadFile(manifestPath)
	if readErr != nil {
		return Manifest{}, readErr
	}
	manifest := Manifest{}
	unmarshalErr := json.Unmarshal(manifestContents, &manifest)
	if unmarshalErr != nil {
		return Manifest{}, unmarshalErr
	}
	return manifest, nil
}

func (r *FileSystemRegistry) getPluginSubdirectoryPath(pluginID ID) (string, error) {
	if validErr := pluginID.Validate(); validErr != nil {
		return "", validErr
	}
	subDirPath := path.Join(r.PluginsDirPath, string(pluginID.Type), pluginID.Name, pluginID.Version)
	return subDirPath, nil
}

func (r *FileSystemRegistry) getPluginManifestPath(pluginID ID) (string, error) {
	pluginPath, subDirErr := r.getPluginSubdirectoryPath(pluginID)
	if subDirErr != nil {
		return "", subDirErr
	}
	manifestPath := path.Join(pluginPath, "MANIFEST.json")
	return manifestPath, nil
}
