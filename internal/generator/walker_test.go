package generator

import (
	"fmt"
	"gorepogen/internal/testutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	asrt := assert.New(t)
	cd, _ := os.Getwd()

	t.Run("entity found", func(t *testing.T) {

		filePath := testutil.CreateTempFile(t, cd, "package entity\n\ntype User struct {}")

		expected := &entityInfo{
			Name:            "User",
			Package:         "entity",
			FullPackagePath: "generator",
		}
		actual, err := Search(cd, "User")

		asrt.NoError(err)
		asrt.Equal(expected, actual)

		testutil.DeleteTempFile(t, filePath)
	})

	t.Run("entity not found", func(t *testing.T) {

		actual, err := Search(cd, "User")

		asrt.Nil(actual)
		asrt.EqualError(err, "can't find given entity: User")
	})

	t.Run("entity not found", func(t *testing.T) {

		dir := "dir/does/not/exists"
		actual, err := Search(dir, "User")

		asrt.Nil(actual)
		asrt.EqualError(err, fmt.Sprintf("lstat %s: no such file or directory", dir))
	})

}

func Test_scanWords(t *testing.T) {

	asrt := assert.New(t)
	cd, _ := os.Getwd()

	filePath := filepath.Join(cd, "temp_1337.go")
	actual, err := scanWords(filePath)

	asrt.Nil(actual)
	asrt.Error(err)
}

func Test_isEntity(t *testing.T) {

	asrt := assert.New(t)
	entityName := "User"

	t.Run("entity match", func(t *testing.T) {
		words := []string{"type", "User", "struct"}
		asrt.True(isEntity(words, 1, words[1], entityName))
	})

	t.Run("entity not match", func(t *testing.T) {
		words := []string{"//", " ", "User", "lorem", "ipsum"}
		asrt.False(isEntity(words, 1, words[1], entityName))
	})

	t.Run("entity not match", func(t *testing.T) {
		words := []string{"User", " ", ":", "=", " "}
		asrt.False(isEntity(words, 1, words[0], entityName))
	})

}

func Test_isDirExcluded(t *testing.T) {

	asrt := assert.New(t)

	t.Run("directory are not excluded", func(t *testing.T) {
		asrt.True(!isDirExcluded("lorem/ipsum/dolor/"))
	})

	t.Run("directory is excluded", func(t *testing.T) {
		asrt.False(!isDirExcluded("lorem/ipsum/.git/dolor/"))
	})

}
