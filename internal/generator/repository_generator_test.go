package generator

import (
	"gorepogen/internal/repocomp"
	"gorepogen/internal/testutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignNamesToComponents(t *testing.T) {

	namesRegistry := NamesRegistry{
		EntityName:            "User",
		PackageName:           "entity",
		FullPackageName:       "project/entity",
		EntityNameWithPackage: "entity.User",
		InterfaceName:         "UserRepository",
		StructName:            "userRepository",
		ConstructorName:       "NewUserRepository",
		ReceiveName:           "r userRepository",
		FileName:              "user_repository.go",
		RepositoryPackageName: "repository",
	}

	components := AssignNamesToComponents(namesRegistry)
	assert.NotNil(t, components)

}

func TestGenerate(t *testing.T) {

	asrt := assert.New(t)
	cd, _ := os.Getwd()

	expected := filepath.Join(cd, "repository", "user_repository.go")

	namesRegistry := NamesRegistry{
		EntityName:            "User",
		PackageName:           "entity",
		FullPackageName:       "project/entity",
		EntityNameWithPackage: "entity.User",
		InterfaceName:         "UserRepository",
		StructName:            "userRepository",
		ConstructorName:       "NewUserRepository",
		ReceiveName:           "r userRepository",
		FileName:              "user_repository.go",
		RepositoryPackageName: "repository",
	}

	components := []repocomp.Appender{
		repocomp.NewInterface(namesRegistry.GetInterfaceNames()),
		repocomp.NewStruct(namesRegistry.GetStructNames()),
		repocomp.NewConstructor(namesRegistry.GetConstructorNames()),
		repocomp.NewMethodsList(namesRegistry.GetMethodListNames()),
	}

	actual, err := Generate(components, namesRegistry, cd)
	asrt.Nil(err)
	asrt.Equal(expected, actual)

	testutil.DeleteTempFile(t, filepath.Dir(actual))
}

func TestResolveNamesRegistry(t *testing.T) {

	asrt := assert.New(t)

	expected := NamesRegistry{
		EntityName:            "User",
		PackageName:           "entity",
		FullPackageName:       "project/entity",
		EntityNameWithPackage: "entity.User",
		InterfaceName:         "UserRepository",
		StructName:            "userRepository",
		ConstructorName:       "NewUserRepository",
		ReceiveName:           "r userRepository",
		FileName:              "user_repository.go",
		RepositoryPackageName: "repository",
	}

	actual := CreateNamesRegistry(&entityInfo{
		Name:            "User",
		Package:         "entity",
		FullPackagePath: "project/entity",
	})

	asrt.Equal(expected, actual)

	t.Run("interface names correct", func(_ *testing.T) {
		actualInterfaceName, actualEntityName, actualFullPackageName := actual.GetInterfaceNames()
		asrt.Equal("UserRepository", actualInterfaceName)
		asrt.Equal("User", actualEntityName)
		asrt.Equal("project/entity", actualFullPackageName)
	})

	t.Run("struct names correct", func(_ *testing.T) {
		asrt.Equal("userRepository", actual.GetStructNames())
	})

	t.Run("constructor names correct", func(_ *testing.T) {
		actualConstructorName, actualInterfaceName, actualStructName := actual.GetConstructorNames()
		asrt.Equal("NewUserRepository", actualConstructorName)
		asrt.Equal("UserRepository", actualInterfaceName)
		asrt.Equal("userRepository", actualStructName)
	})

	t.Run("method names correct", func(_ *testing.T) {
		actualReceiveName, actualEntityNameWithPackage := actual.GetMethodListNames()
		asrt.Equal("r userRepository", actualReceiveName)
		asrt.Equal("entity.User", actualEntityNameWithPackage)
	})

}
