package repocomp

import (
	"testing"

	. "github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func Test_methodsList_AppendTo(t *testing.T) {

	actual := NewFile("repository")

	expected := `package repository

func (r userRepository) FindAll() ([]entity.User, error) {
	var entities []entity.User
	err := r.DB.Find(&entities).Error
	return entities, err
}

func (r userRepository) FindById(id uint) (entity.User, error) {
	var entity entity.User
	err := r.DB.First(&entity, id).Error
	return entity, err
}

func (r userRepository) Save(entity entity.User) (entity.User, error) {
	err := r.DB.Create(&entity).Error
	return entity, err
}

func (r userRepository) Update(entity entity.User) error {
	return r.DB.UpdateColumns(&entity).Error
}

func (r userRepository) Delete(entity entity.User) error {
	return r.DB.Delete(&entity).Error
}

func (r userRepository) Count() (uint, error) {
	var count uint
	err := r.DB.Model(&entity.User{}).Count(&count).Error
	return count, err
}
`

	newMethodList := NewMethodsList("r userRepository", "entity.User")
	newMethodList.AppendTo(actual)

	assert.Equal(t, expected, actual.GoString())

}
