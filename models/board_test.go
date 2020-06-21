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

func Test_PlayDomino(t *testing.T) {
	domino1 := &models.Domino{}
	err := domino1.Set(1, 6)
	assert.NoError(t, err)

	domino2 := &models.Domino{}
	err = domino2.Set(2, 6)
	assert.NoError(t, err)

	board, err := models.NewBoard(domino1)
	assert.NoError(t, err)
	assert.NotNil(t, board)

	// Valid at tail - 2:6 works on 6
	err = board.PlayDomino(domino2, false)
	assert.NoError(t, err)
	assert.Equal(t, 1, board.GetHead())
	assert.Equal(t, 2, board.GetTail())
	assert.Equal(t, 2, len(board.GetPlayedDominos()))

	//Invalid at head - 2:6 doesn't works on 1
	err = board.PlayDomino(domino2, true)
	assert.EqualError(t, err, "Invalid Move")
	assert.Equal(t, 1, board.GetHead())
	assert.Equal(t, 2, board.GetTail())
	assert.Equal(t, 2, len(board.GetPlayedDominos()))
}
