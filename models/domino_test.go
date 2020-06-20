package models_test

import (
	"testing"

	"github.com/neiln3121/dominos/models"
	"github.com/stretchr/testify/assert"
)

func Test_DominoSet(t *testing.T) {
	domino := models.Domino{}

	err := domino.Set(-1, 0)
	assert.Error(t, err)

	err = domino.Set(1, 9)
	assert.Error(t, err)
}

func Test_DominoTotal(t *testing.T) {
	domino := models.Domino{}

	err := domino.Set(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, 0, domino.Total())

	err = domino.Set(1, 5)
	assert.NoError(t, err)
	assert.Equal(t, 6, domino.Total())

	err = domino.Set(6, 6)
	assert.NoError(t, err)
	assert.Equal(t, 12, domino.Total())
}
