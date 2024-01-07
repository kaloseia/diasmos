package plugin

func NewFileSystemRegistry(srcDirPath string) (*FileSystemRegistry, error) {
	registry := &FileSystemRegistry{
		PluginsDirPath: srcDirPath,
	}
	initErr := registry.Initialize()
	if initErr != nil {
		return nil, initErr
	}
	return registry, nil
}
