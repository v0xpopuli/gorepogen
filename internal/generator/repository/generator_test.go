package repository

import (
	"os"
	"path/filepath"
	"testing"

	g "github.com/v0xpopuli/gorepogen/internal/generator"
	"github.com/v0xpopuli/gorepogen/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {

	asrt := assert.New(t)
	cd, _ := os.Getwd()

	expected := filepath.Join(cd, "repository", "user_repository.go")

	namesRegistry := g.NamesRegister{
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

	actual, err := NewGenerator(namesRegistry).Generate(cd)
	asrt.Nil(err)
	asrt.Equal(expected, actual)

	testutil.DeleteTempFile(t, filepath.Dir(actual))
}

func TestResolveNamesRegistry(t *testing.T) {

	asrt := assert.New(t)

	expected := g.NamesRegister{
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

	actual := g.NewNamesRegister(&g.EntityInfo{
		EntityName:      "User",
		EntityPackage:   "entity",
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
