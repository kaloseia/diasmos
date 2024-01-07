package plugin_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/diasmos/domain/plugin"
	"github.com/kaloseia/diasmos/domain/testutils"
)

type IDTestSuite struct {
	suite.Suite

	TestDirPath string
}

func TestIDTestSuite(t *testing.T) {
	suite.Run(t, new(IDTestSuite))
}

func (suite *IDTestSuite) SetupTest() {
	suite.TestDirPath = testutils.GetTestDirPath()
}

func (suite *IDTestSuite) TearDownTest() {
	suite.TestDirPath = ""
}

func (suite *IDTestSuite) TestFromKey() {
	id := plugin.ID{}

	keyErr := id.FromKey("compile/go-struct@5.3.2")

	suite.Nil(keyErr)
	suite.Equal(id.Name, "go-struct")
	suite.Equal(id.Version, "5.3.2")
	suite.Equal(id.Type, plugin.TypeCompile)
}

func (suite *IDTestSuite) TestFromKey_InvalidFormat() {
	id := plugin.ID{}

	keyErr := id.FromKey("compile+go-struct@5.3.2")

	suite.ErrorIs(keyErr, plugin.ErrKeyInvalidFormat)
	suite.Zero(id)
}

func (suite *IDTestSuite) TestFromKey_InvalidFormat_Empty() {
	id := plugin.ID{}

	keyErr := id.FromKey("")

	suite.ErrorIs(keyErr, plugin.ErrKeyInvalidFormat)
	suite.Zero(id)
}

func (suite *IDTestSuite) TestFromKey_InvalidFormat_Type() {
	id := plugin.ID{}

	keyErr := id.FromKey("invalid/go-struct@5.3.2")

	suite.ErrorIs(keyErr, plugin.ErrKeyInvalidFormat)
	suite.Zero(id)
}

func (suite *IDTestSuite) TestFromKey_InvalidFormat_Name() {
	id := plugin.ID{}

	keyErr := id.FromKey("compile/go-struct$#!@5.3.2")

	suite.ErrorIs(keyErr, plugin.ErrKeyInvalidFormat)
	suite.Zero(id)
}

func (suite *IDTestSuite) TestFromKey_InvalidFormat_Version() {
	id := plugin.ID{}

	keyErr := id.FromKey("compile/go-struct@5.3.2.X")

	suite.ErrorIs(keyErr, plugin.ErrKeyInvalidFormat)
	suite.Zero(id)
}

func (suite *IDTestSuite) TestToKey() {
	id := plugin.ID{
		Name:    "go-struct",
		Version: "5.3.2",
		Type:    plugin.TypeCompile,
	}

	key, keyErr := id.ToKey()

	suite.Nil(keyErr)
	suite.Equal(key, "compile/go-struct@5.3.2")

	suite.Equal(id.Name, "go-struct")
	suite.Equal(id.Version, "5.3.2")
	suite.Equal(id.Type, plugin.TypeCompile)
}

func (suite *IDTestSuite) TestToKey_MissingField_Type() {
	id := plugin.ID{
		Name:    "go-struct",
		Version: "5.3.2",
		Type:    "",
	}

	key, keyErr := id.ToKey()

	suite.ErrorIs(keyErr, plugin.ErrIDInvalidFieldType)
	suite.Equal(key, "")

	suite.Equal(id.Name, "go-struct")
	suite.Equal(id.Version, "5.3.2")
	suite.EqualValues(id.Type, "")
}

func (suite *IDTestSuite) TestToKey_MissingField_Name() {
	id := plugin.ID{
		Name:    "",
		Version: "5.3.2",
		Type:    plugin.TypeCompile,
	}

	key, keyErr := id.ToKey()

	suite.ErrorIs(keyErr, plugin.ErrIDInvalidFieldName)
	suite.Equal(key, "")

	suite.Equal(id.Name, "")
	suite.Equal(id.Version, "5.3.2")
	suite.Equal(id.Type, plugin.TypeCompile)
}

func (suite *IDTestSuite) TestToKey_MissingField_Version() {
	id := plugin.ID{
		Name:    "go-struct",
		Version: "",
		Type:    plugin.TypeCompile,
	}

	key, keyErr := id.ToKey()

	suite.ErrorIs(keyErr, plugin.ErrIDInvalidFieldVersion)
	suite.Equal(key, "")

	suite.Equal(id.Name, "go-struct")
	suite.Equal(id.Version, "")
	suite.Equal(id.Type, plugin.TypeCompile)
}

func (suite *IDTestSuite) TestToKey_InvalidField_Type() {
	id := plugin.ID{
		Name:    "go-struct",
		Version: "5.3.2",
		Type:    "",
	}

	key, keyErr := id.ToKey()

	suite.ErrorIs(keyErr, plugin.ErrIDInvalidFieldType)
	suite.Equal(key, "")

	suite.Equal(id.Name, "go-struct")
	suite.Equal(id.Version, "5.3.2")
	suite.EqualValues(id.Type, "")
}

func (suite *IDTestSuite) TestToKey_InvalidField_Name() {
	id := plugin.ID{
		Name:    ";DROP TABLE",
		Version: "5.3.2",
		Type:    plugin.TypeCompile,
	}

	key, keyErr := id.ToKey()

	suite.ErrorIs(keyErr, plugin.ErrIDInvalidFieldName)
	suite.Equal(key, "")

	suite.Equal(id.Name, ";DROP TABLE")
	suite.Equal(id.Version, "5.3.2")
	suite.Equal(id.Type, plugin.TypeCompile)
}

func (suite *IDTestSuite) TestToKey_InvalidField_Version() {
	id := plugin.ID{
		Name:    "go-struct",
		Version: "x123.432.123",
		Type:    plugin.TypeCompile,
	}

	key, keyErr := id.ToKey()

	suite.ErrorIs(keyErr, plugin.ErrIDInvalidFieldVersion)
	suite.Equal(key, "")

	suite.Equal(id.Name, "go-struct")
	suite.Equal(id.Version, "x123.432.123")
	suite.Equal(id.Type, plugin.TypeCompile)
}
