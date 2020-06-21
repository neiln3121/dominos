package models

import (
	"errors"
)

// MAXDOMINOS - maximum number of dominos a player can have
const MAXDOMINOS int = 7

// Player - stores detaisl for a game participant
type Player struct {
	ID      int
	dominos []*Domino
}

// AddDomino - add a domino to a player
func (p *Player) AddDomino(domino *Domino) error {
	if domino == nil {
		return errors.New("Not a valid domino")
	}
	p.dominos = append(p.dominos, domino)

	return nil
}

// RemoveDomino - remove a domino to a player
func (p *Player) RemoveDomino(index int) error {
	if !p.HasDomino(index) {
		return errors.New("Player does not have that domino")
	}
	p.dominos = append(p.dominos[:index], p.dominos[index+1:]...)

	return nil
}

func (p *Player) GetHighestDouble() (highest, index int) {
	highest = -1
	for i, domino := range p.dominos {
		// Must be a double
		if domino.half[0] == domino.half[1] {
			if domino.half[0] > highest {
				highest = domino.half[0]
				index = i
			}
		}
	}
	return
}

func (p *Player) CanProceed(head, tail int) bool {
	for _, domino := range p.GetDominos() {
		if domino.half[0] == head || domino.half[0] == tail ||
			domino.half[1] == head || domino.half[1] == tail {
			return true
		}
	}
	return false
}

func (p *Player) Get(index int) (*Domino, error) {
	if !p.HasDomino(index) {
		return nil, errors.New("Player does not have that domino")
	}

	return p.dominos[index], nil
}

func (p *Player) GetDominos() []*Domino {
	return p.dominos
}

func (p *Player) HasStartingDominos() bool {
	return p.DominoCount() == MAXDOMINOS
}

func (p *Player) DominoCount() int {
	return len(p.dominos)
}

func (p *Player) TotalDots() int {
	sum := 0
	for _, domino := range p.dominos {
		sum += domino.Total()
	}
	return sum
}

func (p *Player) HasDomino(index int) bool {
	if index < 0 || index >= p.DominoCount() {
		return false
	}
	return true
}

// NewPlayer - new instance
func NewPlayer(index int) *Player {
	dominos := new([]*Domino)
	return &Player{ID: index + 1, dominos: *dominos}
}
