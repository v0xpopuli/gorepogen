package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var entityName = "User"

func Test_countFuncBody(t *testing.T) {

	expected := `var count uint
err := r.DB.Model(&User{}).Count(&count).Error
return count, err`

	actual := countFuncBody(entityName)

	assert.Equal(t, expected, actual.GoString())
}

func Test_deleteFuncBody(t *testing.T) {

	expected := `return r.DB.Delete(&entity).Error`

	actual := deleteFuncBody()

	assert.Equal(t, expected, actual.GoString())
}

func Test_findAllFuncBody(t *testing.T) {

	expected := `var entities []User
err := r.DB.Find(&entities).Error
return entities, err`

	actual := findAllFuncBody(entityName)

	assert.Equal(t, expected, actual.GoString())
}

func Test_findByIdFuncBody(t *testing.T) {

	expected := `var entity User
err := r.DB.First(&entity, id).Error
return entity, err`

	actual := findByIdFuncBody(entityName)

	assert.Equal(t, expected, actual.GoString())
}

func Test_saveFuncBody(t *testing.T) {

	expected := `err := r.DB.Create(&entity).Error
return entity, err`

	actual := saveFuncBody()

	assert.Equal(t, expected, actual.GoString())
}

func Test_updateFuncBody(t *testing.T) {

	expected := `return r.DB.UpdateColumns(&entity).Error`

	actual := updateFuncBody()

	assert.Equal(t, expected, actual.GoString())
}
