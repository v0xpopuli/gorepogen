package component

import (
	"testing"

	. "github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func Test_constructorGenerator_AppendTo(t *testing.T) {

	actual := NewFile("repository")

	expected := `package repository

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}
`

	newConstructor := NewConstructor(
		"NewUserRepository",
		"UserRepository",
		"userRepository",
	)
	newConstructor.AppendTo(actual)

	assert.Equal(t, expected, actual.GoString())
}
