package models

import (
	"errors"
)

// DOMINOS - maximum number of dominos in the game
const DOMINOS int = 28

// Table - stores the dominos ready to be picked from by a player
type Table struct {
	Dominos     []Domino
	pickedCount int
}

// GetUnpickedDomino - get a previously unpicked domino
func (t *Table) GetUnpickedDomino(index int) (*Domino, error) {
	if index < 0 || index >= DOMINOS {
		return nil, errors.New("Invalid domino ID")
	}

	domino := &t.Dominos[index]
	if domino.IsPicked == true {
		return nil, errors.New("This domino has already been picked")
	}
	domino.IsPicked = true
	t.pickedCount++

	return domino, nil
}

// GetDominos -returns all the dominos on the table
func (t *Table) GetDominos() []Domino {
	return t.Dominos
}

// AllPicked - returns whether every domino has been picked
func (t *Table) AllPicked() bool {
	return t.pickedCount == DOMINOS
}

// NewTable - new instance
func NewTable() *Table {
	var table Table
	table.Dominos = make([]Domino, DOMINOS)

	index := 0
	for i := 0; i < 7; i++ {
		for j := i; j < 7; j++ {
			table.Dominos[index].Set(i, j)
			index++
		}
	}

	return &table
}
