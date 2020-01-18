package repocomp

import (
	"testing"

	. "github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func Test_structGenerator_AppendTo(t *testing.T) {

	actual := NewFile("repository")

	expected := `package repository

import gorm "github.com/jinzhu/gorm"

type userRepository struct {
	*gorm.DB
}
`

	newStruct := NewStruct("userRepository")
	newStruct.AppendTo(actual)

	assert.Equal(t, expected, actual.GoString())

}
