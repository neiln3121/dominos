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

	player := models.NewPlayer()

	player.AddDomino(domino)
	assert.Equal(t, 1, player.DominoCount())
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	player.AddDomino(domino)
	assert.Equal(t, 7, player.DominoCount())
	assert.True(t, player.HasStartingDominos())

	err = player.AddDomino(nil)
	assert.Error(t, err)

	player.AddDomino(domino)
	assert.Equal(t, 8, player.DominoCount())
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

	player := models.NewPlayer()
	player.AddDomino(domino1)
	player.AddDomino(domino2)
	player.AddDomino(domino3)
	assert.Equal(t, 3, player.DominoCount())

	err = player.RemoveDomino(10)
	assert.Error(t, err)

	err = player.RemoveDomino(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, player.DominoCount())

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

	player := models.NewPlayer()
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

func Test_CanProceed(t *testing.T) {
	domino1 := &models.Domino{}
	err := domino1.Set(1, 1)
	assert.NoError(t, err)

	domino2 := &models.Domino{}
	err = domino2.Set(6, 5)
	assert.NoError(t, err)

	domino3 := &models.Domino{}
	err = domino3.Set(4, 4)
	assert.NoError(t, err)

	player := models.NewPlayer()
	player.AddDomino(domino1)
	player.AddDomino(domino2)
	player.AddDomino(domino3)

	assert.True(t, player.CanProceed(1, 0))
	assert.True(t, player.CanProceed(4, 0))
	assert.True(t, player.CanProceed(4, 4))
	assert.True(t, player.CanProceed(5, 4))

	assert.False(t, player.CanProceed(2, 3))
	assert.False(t, player.CanProceed(2, 0))
	assert.False(t, player.CanProceed(0, 0))
}

func Test_TotalDots(t *testing.T) {
	domino1 := &models.Domino{}
	err := domino1.Set(1, 1)
	assert.NoError(t, err)

	domino2 := &models.Domino{}
	err = domino2.Set(6, 5)
	assert.NoError(t, err)

	domino3 := &models.Domino{}
	err = domino3.Set(4, 4)
	assert.NoError(t, err)

	player := models.NewPlayer()
	player.AddDomino(domino1)
	assert.Equal(t, 2, player.TotalDots())

	player.AddDomino(domino2)
	assert.Equal(t, 13, player.TotalDots())

	player.AddDomino(domino3)
	assert.Equal(t, 21, player.TotalDots())

	err = player.RemoveDomino(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, player.DominoCount())
	assert.Equal(t, 10, player.TotalDots())
}
