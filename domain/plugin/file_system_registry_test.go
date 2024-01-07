package plugin_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/diasmos/domain/plugin"
	"github.com/kaloseia/diasmos/domain/testutils"
)

type FileSystemRegistryTestSuite struct {
	suite.Suite

	TestDirPath   string
	PluginDirPath string
	TempDirPath   string
}

func TestFileSystemRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(FileSystemRegistryTestSuite))
}

func (suite *FileSystemRegistryTestSuite) SetupTest() {
	suite.TestDirPath = testutils.GetTestDirPath()

	suite.PluginDirPath = filepath.Join(suite.TestDirPath, "registry", "plugins")
	suite.TempDirPath = filepath.Join(suite.TestDirPath, "registry", "tmp")
}

func (suite *FileSystemRegistryTestSuite) TearDownTest() {
	suite.TestDirPath = ""
}

func (suite *FileSystemRegistryTestSuite) TestInitialize() {
	registry := plugin.FileSystemRegistry{
		PluginsDirPath: suite.TestDirPath + "/plugins",
	}

	initErr := registry.Initialize()

	suite.Nil(initErr)
	suite.True(registry.IsInitialized())
}

func (suite *FileSystemRegistryTestSuite) TestInitialize_AlreadyInitialized() {
	registry := plugin.FileSystemRegistry{
		PluginsDirPath: suite.TestDirPath + "/plugins",
	}

	initErr0 := registry.Initialize()

	suite.Nil(initErr0)
	suite.True(registry.IsInitialized())

	initErr1 := registry.Initialize()

	suite.ErrorIs(initErr1, plugin.ErrAlreadyInitialized)
	suite.True(registry.IsInitialized())
}

func (suite *FileSystemRegistryTestSuite) TestInitialize_PathNotSet() {
	registry := plugin.FileSystemRegistry{}

	initErr := registry.Initialize()

	suite.ErrorIs(initErr, plugin.ErrNoPluginsDirPath)
	suite.False(registry.IsInitialized())
}

func (suite *FileSystemRegistryTestSuite) TestValidatePlugin() {
	registry, registryErr := plugin.NewFileSystemRegistry(suite.TestDirPath + "/registry/plugins")
	suite.Nil(registryErr)

	id := plugin.ID{
		Name:    "go-struct",
		Version: "0.0.0",
		Type:    plugin.TypeCompile,
	}

	manifest, initErr := registry.ValidatePlugin(id)

	suite.Nil(initErr)
	suite.Equal(manifest.ID.Name, "go-struct")
	suite.Equal(manifest.ID.Version, "0.0.0")
	suite.Equal(manifest.ID.Type, plugin.TypeCompile)
	suite.Equal(manifest.Author, "John Doe")
	suite.Equal(manifest.Description, "The go-struct plugin is a compile plugin for transforming Dia structures into basic go structures with the correct fields and data types.")
}
