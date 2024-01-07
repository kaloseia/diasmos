package plugin_test

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/diasmos/domain/files"
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
	registryPath := path.Join(suite.TestDirPath, "/registry/plugins")
	registry, registryErr := plugin.NewFileSystemRegistry(registryPath)
	suite.Nil(registryErr)

	id := plugin.ID{
		Name:    "go-struct",
		Version: "0.0.0",
		Type:    plugin.TypeCompile,
	}

	manifest, validateErr := registry.ValidatePlugin(id)

	suite.Nil(validateErr)
	suite.Equal(manifest.ID.Name, "go-struct")
	suite.Equal(manifest.ID.Version, "0.0.0")
	suite.Equal(manifest.ID.Type, plugin.TypeCompile)
	suite.Equal(manifest.Author, "John Doe")
	suite.Equal(manifest.Description, "The go-struct plugin is a compile plugin for transforming Dia structures into basic go structures with the correct fields and data types.")
}

func (suite *FileSystemRegistryTestSuite) TestValidatePlugin_NotFound_Name() {
	registryPath := path.Join(suite.TestDirPath, "/registry/plugins")
	registry, registryErr := plugin.NewFileSystemRegistry(registryPath)
	suite.Nil(registryErr)

	id := plugin.ID{
		Name:    "INVALID",
		Version: "0.0.0",
		Type:    plugin.TypeCompile,
	}

	manifest, validateErr := registry.ValidatePlugin(id)

	suite.ErrorContains(validateErr, "compile/INVALID@0.0.0")
	suite.ErrorContains(validateErr, "manifest not found")
	suite.Zero(manifest)
}

func (suite *FileSystemRegistryTestSuite) TestValidatePlugin_NotFound_Version() {
	registryPath := path.Join(suite.TestDirPath, "/registry/plugins")
	registry, registryErr := plugin.NewFileSystemRegistry(registryPath)
	suite.Nil(registryErr)

	id := plugin.ID{
		Name:    "go-struct",
		Version: "99999.0.0",
		Type:    plugin.TypeCompile,
	}

	manifest, validateErr := registry.ValidatePlugin(id)

	suite.ErrorContains(validateErr, "compile/go-struct@99999.0.0")
	suite.ErrorContains(validateErr, "manifest not found")
	suite.Zero(manifest)
}

func (suite *FileSystemRegistryTestSuite) TestValidatePlugin_NotFound_Type() {
	registryPath := path.Join(suite.TestDirPath, "/registry/plugins")
	registry, registryErr := plugin.NewFileSystemRegistry(registryPath)
	suite.Nil(registryErr)

	id := plugin.ID{
		Name:    "go-struct",
		Version: "0.0.0",
		Type:    plugin.TypeExecution,
	}

	manifest, validateErr := registry.ValidatePlugin(id)

	suite.ErrorContains(validateErr, "execute/go-struct@0.0.0")
	suite.ErrorContains(validateErr, "manifest not found")
	suite.Zero(manifest)
}
