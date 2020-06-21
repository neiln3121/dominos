package models_test

import (
	"testing"

	"github.com/neiln3121/dominos/models"
	"github.com/stretchr/testify/assert"
)

func Test_NewBoard_Valid(t *testing.T) {
	domino := &models.Domino{}

	err := domino.Set(1, 6)
	assert.NoError(t, err)

	board, err := models.NewBoard(domino)
	assert.NoError(t, err)
	assert.NotNil(t, board)
	assert.Equal(t, 1, board.GetHead())
	assert.Equal(t, 6, board.GetTail())
	assert.Equal(t, 1, len(board.GetPlayedDominos()))
}

func Test_NewBoard_Invalid(t *testing.T) {
	board, err := models.NewBoard(nil)
	assert.Error(t, err)
	assert.Nil(t, board)
}
