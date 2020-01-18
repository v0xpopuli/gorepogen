package repocomp

import (
	"testing"

	. "github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func Test_interfaceGenerator_AppendTo(t *testing.T) {

	actual := NewFile("repository")

	expected := `package repository

import entity "app/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindById(uint) (entity.User, error)
	Save(entity.User) (entity.User, error)
	Update(entity.User) error
	Delete(entity.User) error
	Count() (uint, error)
}
`

	newInterface := NewInterface(
		"UserRepository",
		"User",
		"app/entity",
	)
	newInterface.AppendTo(actual)

	assert.Equal(t, expected, actual.GoString())
}
