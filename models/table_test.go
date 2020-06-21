package models_test

import (
	"testing"

	"github.com/neiln3121/dominos/models"
	"github.com/stretchr/testify/assert"
)

func Test_AllPicked(t *testing.T) {
	table := models.NewTable()

	assert.False(t, table.AllPicked())

	dominoCount := len(table.GetDominos())
	assert.Equal(t, 28, dominoCount)

	var err error
	for i := 0; i < dominoCount; i++ {
		_, err = table.GetUnpickedDomino(i)
		assert.NoError(t, err)
	}

	assert.True(t, table.AllPicked())
}

func Test_GetUnpickedDomino(t *testing.T) {
	table := models.NewTable()

	assert.False(t, table.AllPicked())

	dominoCount := len(table.GetDominos())
	assert.Equal(t, 28, dominoCount)

	domino, err := table.GetUnpickedDomino(0)
	assert.NoError(t, err)
	assert.NotNil(t, domino)

	domino, err = table.GetUnpickedDomino(27)
	assert.NoError(t, err)
	assert.NotNil(t, domino)

	domino, err = table.GetUnpickedDomino(-1)
	assert.Error(t, err)
	assert.Nil(t, domino)

	domino, err = table.GetUnpickedDomino(28)
	assert.Error(t, err)
	assert.Nil(t, domino)
}
