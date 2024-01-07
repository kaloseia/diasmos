package plugin_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/diasmos/domain/plugin"
	"github.com/kaloseia/diasmos/domain/testutils"
)

type ManifestTestSuite struct {
	suite.Suite

	TestDirPath string
}

func TestManifestTestSuite(t *testing.T) {
	suite.Run(t, new(ManifestTestSuite))
}

func (suite *ManifestTestSuite) SetupTest() {
	suite.TestDirPath = testutils.GetTestDirPath()
}

func (suite *ManifestTestSuite) TearDownTest() {
	suite.TestDirPath = ""
}

func (suite *ManifestTestSuite) TestUnmarshalJSON() {
	jsonData := `{
		"name": "go-gorm",
		"version": "0.0.0",
		"type": "compile",
		"author": "John Doe",
		"description": "A description here.",
		"dependencies": {
			"compile/go-struct": "1.0.4"
		}
	}`

	var manifest plugin.Manifest
	unmarshalErr := json.Unmarshal([]byte(jsonData), &manifest)

	suite.Nil(unmarshalErr)

	suite.Equal(manifest.ID.Name, "go-gorm")
	suite.Equal(manifest.ID.Version, "0.0.0")
	suite.Equal(manifest.ID.Type, plugin.TypeCompile)
	suite.Equal(manifest.Author, "John Doe")
	suite.Equal(manifest.Description, "A description here.")
}
