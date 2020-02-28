package walker

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/v0xpopuli/gorepogen/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	asrt := assert.New(t)
	cd, _ := os.Getwd()

	walker := NewWalker(
		filepath.Base(cd),
		"User",
	)

	t.Run("entity found", func(t *testing.T) {

		filePath := testutil.CreateTempFile(t, cd, "package entity\n\ntype User struct {}")

		expected := &EntityInfo{
			EntityName:      "User",
			EntityPackage:   "entity",
			FullPackagePath: "generator",
		}
		actual, err := walker.Walk(cd)

		asrt.NoError(err)
		asrt.Equal(expected, actual)

		testutil.DeleteTempFile(t, filePath)
	})

	t.Run("entity not found", func(t *testing.T) {

		actual, err := walker.Walk(cd)

		asrt.Nil(actual)
		asrt.EqualError(err, "Can't find given entity: User")
	})

	t.Run("entity not found", func(t *testing.T) {

		dir := "dir/does/not/exists"
		actual, err := walker.Walk(dir)

		asrt.Nil(actual)
		asrt.EqualError(err, fmt.Sprintf("lstat %s: no such file or directory", dir))
	})

}

func Test_isEntity(t *testing.T) {

	asrt := assert.New(t)
	entityName := "User"

	walker := NewWalker(
		"",
		entityName,
	)

	t.Run("entity match", func(t *testing.T) {
		asrt.True(walker.isEntity("package entity\n\ntype User struct {}"))
	})

	t.Run("entity not match", func(t *testing.T) {
		asrt.False(walker.isEntity("package entity\n\ntype Person struct {}"))
	})

}

func Test_isDirExcluded(t *testing.T) {

	asrt := assert.New(t)

	walker := NewWalker(
		"",
		"",
	)

	t.Run("directory are not excluded", func(t *testing.T) {
		asrt.True(!walker.isDirExcluded("lorem/ipsum/dolor/"))
	})

	t.Run("directory is excluded", func(t *testing.T) {
		asrt.False(!walker.isDirExcluded("lorem/ipsum/.git/dolor/"))
	})

}
