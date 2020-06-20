package models_test

import (
	"testing"

	"github.com/neiln3121/dominos/models"
	"github.com/stretchr/testify/assert"
)

func Test_AddDomino(t *testing.T) {
	domino := &models.Domino{}
	err := domino.Set(1, 1)
	assert.NoError(t, err)

	player := models.NewPlayer(1)

	player.AddDomino(domino)
	assert.Equal(t, 1, player.DominoCount())
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	assert.Equal(t, 7, player.DominoCount())

	err = player.AddDomino(domino)
	assert.Equal(t, 7, player.DominoCount())
	assert.Error(t, err)
}

func Test_RemoveDomino(t *testing.T) {
	domino1 := &models.Domino{}
	err := domino1.Set(1, 1)
	assert.NoError(t, err)

	domino2 := &models.Domino{}
	err = domino2.Set(6, 5)
	assert.NoError(t, err)

	domino3 := &models.Domino{}
	err = domino3.Set(4, 4)
	assert.NoError(t, err)

	player := models.NewPlayer(1)
	player.AddDomino(domino1)
	player.AddDomino(domino2)
	player.AddDomino(domino3)

	assert.Equal(t, 3, player.DominoCount())
	assert.Equal(t, 21, player.TotalDots())

	err = player.RemoveDomino(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, player.DominoCount())
	assert.Equal(t, 10, player.TotalDots())

	err = player.RemoveDomino(2)
	assert.Error(t, err)
}

func Test_GetHighestDouble(t *testing.T) {
	domino1 := &models.Domino{}
	err := domino1.Set(1, 1)
	assert.NoError(t, err)

	domino2 := &models.Domino{}
	err = domino2.Set(6, 5)
	assert.NoError(t, err)

	domino3 := &models.Domino{}
	err = domino3.Set(4, 4)
	assert.NoError(t, err)

	player := models.NewPlayer(1)
	player.AddDomino(domino1)
	player.AddDomino(domino2)
	player.AddDomino(domino3)

	highest, index := player.GetHighestDouble()
	highestDomino, err := player.Get(index)
	assert.NoError(t, err)

	assert.Equal(t, 4, highest)
	assert.Equal(t, 8, highestDomino.Total())

	assert.Equal(t, 21, player.TotalDots())
}
