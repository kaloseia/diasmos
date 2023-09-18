package dia_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/kaloseia/diasmos/domain/dia"
	"github.com/kaloseia/diasmos/domain/testutils"
)

type RegistryTestSuite struct {
	suite.Suite

	TestDirPath string

	ModelsDirPath   string
	EntitiesDirPath string
}

func TestRegistryTestSuite(t *testing.T) {
	suite.Run(t, new(RegistryTestSuite))
}

func (suite *RegistryTestSuite) SetupTest() {
	suite.TestDirPath = testutils.GetTestDirPath()

	suite.ModelsDirPath = filepath.Join(suite.TestDirPath, "registry", "models")
	suite.EntitiesDirPath = filepath.Join(suite.TestDirPath, "registry", "entities")
}

func (suite *RegistryTestSuite) TearDownTest() {
	suite.TestDirPath = ""
	dia.ResetRegistry()
}

func (suite *RegistryTestSuite) TestLoadModelsFromDirectory() {
	registry := dia.GetRegistry()

	modelsErr := registry.LoadModelsFromDirectory(suite.ModelsDirPath)

	suite.Nil(modelsErr)
	suite.Len(registry.Models, 3)
	suite.Len(registry.Entities, 0)

	model0, modelExists0 := registry.Models["Company"]
	suite.True(modelExists0)
	suite.Equal(model0.Name, "Company")

	suite.Len(model0.Fields, 4)

	modelField00, fieldExists00 := model0.Fields["UUID"]
	suite.True(fieldExists00)
	suite.Equal(modelField00.Type, "UUID")
	suite.Len(modelField00.Attributes, 2)
	suite.Equal(modelField00.Attributes[0], "immutable")
	suite.Equal(modelField00.Attributes[1], "mandatory")

	modelField01, fieldExists01 := model0.Fields["ID"]
	suite.True(fieldExists01)
	suite.Equal(modelField01.Type, "AutoIncrement")
	suite.Len(modelField01.Attributes, 1)
	suite.Equal(modelField01.Attributes[0], "mandatory")

	modelField02, fieldExists02 := model0.Fields["FoundedAt"]
	suite.True(fieldExists02)
	suite.Equal(modelField02.Type, "Time")
	suite.Len(modelField02.Attributes, 0)

	modelField03, fieldExists03 := model0.Fields["Name"]
	suite.True(fieldExists03)
	suite.Equal(modelField03.Type, "String")
	suite.Len(modelField03.Attributes, 0)

	suite.Len(model0.Identifiers, 2)

	modelIDs00, idsExist00 := model0.Identifiers["primary"]
	suite.True(idsExist00)
	suite.ElementsMatch(modelIDs00.Fields, []string{"UUID"})

	modelIDs01, idsExist01 := model0.Identifiers["record"]
	suite.True(idsExist01)
	suite.ElementsMatch(modelIDs01.Fields, []string{"ID"})

	suite.Len(model0.Related, 2)

	modelRelated00, relatedExists00 := model0.Related["Address"]
	suite.True(relatedExists00)
	suite.Equal(modelRelated00.Type, "HasOne")

	modelRelated01, relatedExists01 := model0.Related["ContactInfo"]
	suite.True(relatedExists01)
	suite.Equal(modelRelated01.Type, "HasOne")

	model1, modelExists1 := registry.Models["Address"]
	suite.True(modelExists1)
	suite.Equal(model1.Name, "Address")

	suite.Len(model1.Fields, 3)

	modelField10, fieldExists10 := model1.Fields["ID"]
	suite.True(fieldExists10)
	suite.Equal(modelField10.Type, "AutoIncrement")
	suite.Len(modelField10.Attributes, 1)
	suite.Equal(modelField10.Attributes[0], "mandatory")

	modelField11, fieldExists11 := model1.Fields["Street"]
	suite.True(fieldExists11)
	suite.Equal(modelField11.Type, "String")
	suite.Len(modelField11.Attributes, 0)

	modelField12, fieldExists12 := model1.Fields["HouseNumber"]
	suite.True(fieldExists12)
	suite.Equal(modelField12.Type, "String")
	suite.Len(modelField12.Attributes, 0)

	suite.Len(model1.Identifiers, 2)
	modelIDs10, idsExist10 := model1.Identifiers["primary"]
	suite.True(idsExist10)
	suite.ElementsMatch(modelIDs10.Fields, []string{"ID"})

	modelIDs11, idsExist11 := model1.Identifiers["street"]
	suite.True(idsExist11)
	suite.ElementsMatch(modelIDs11.Fields, []string{"Street", "HouseNumber"})

	suite.Len(model1.Related, 1)

	modelRelated10, relatedExists10 := model1.Related["Company"]
	suite.True(relatedExists10)
	suite.Equal(modelRelated10.Type, "ForOne")

	model2, modelExists2 := registry.Models["ContactInfo"]
	suite.True(modelExists2)
	suite.Equal(model2.Name, "ContactInfo")

	suite.Len(model2.Fields, 3)

	modelField20, fieldExists20 := model2.Fields["ID"]
	suite.True(fieldExists20)
	suite.Equal(modelField20.Type, "AutoIncrement")
	suite.Len(modelField20.Attributes, 1)
	suite.Equal(modelField20.Attributes[0], "mandatory")

	modelField21, fieldExists21 := model2.Fields["Email"]
	suite.True(fieldExists21)
	suite.Equal(modelField21.Type, "String")
	suite.Len(modelField21.Attributes, 1)
	suite.Equal(modelField21.Attributes[0], "mandatory")

	modelField22, fieldExists22 := model2.Fields["PhoneNumber"]
	suite.True(fieldExists22)
	suite.Equal(modelField22.Type, "String")
	suite.Len(modelField22.Attributes, 0)

	suite.Len(model2.Identifiers, 1)
	modelID20, idExists20 := model2.Identifiers["primary"]
	suite.True(idExists20)
	suite.ElementsMatch(modelID20.Fields, []string{"ID"})

	suite.Len(model2.Related, 1)

	modelRelated20, relatedExists20 := model2.Related["Company"]
	suite.True(relatedExists20)
	suite.Equal(modelRelated20.Type, "ForOne")
}
