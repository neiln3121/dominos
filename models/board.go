package models

import (
	"errors"
)

// Board - stores all the played dominos and history
type Board struct {
	played []*Domino
	head   int
	tail   int
}

// PlayDomino - play a domino onto the board at either the head or the tail of the board
func (b *Board) PlayDomino(domino *Domino, atHead bool) error {
	var result int
	if atHead {
		result = b.checkDominoPlayableHead(domino)
		if result != -1 {
			b.played = append([]*Domino{domino}, b.played...)
			b.head = result
		}
	} else {
		result = b.checkDominoPlayableTail(domino)
		if result != -1 {
			b.played = append(b.played, domino)
			b.tail = result
		}
	}

	if result == -1 {
		return errors.New("Invalid Move")
	}
	return nil
}

// GetPlayedDominos - shows all the dominos played on the board
func (b *Board) GetPlayedDominos() []*Domino {
	return b.played
}

// GetHead - get the start or the head of the domino trail
func (b *Board) GetHead() int {
	return b.head
}

// GetTail - get the end or the tail of the domino trail
func (b *Board) GetTail() int {
	return b.tail
}

func (b *Board) checkDominoPlayableHead(domino *Domino) int {
	if domino.half[0] == b.head {
		domino.playedFlipped = true
		return domino.half[1]
	} else if domino.half[1] == b.head {
		return domino.half[0]
	}
	return -1
}

func (b *Board) checkDominoPlayableTail(domino *Domino) int {
	if domino.half[0] == b.tail {
		return domino.half[1]
	} else if domino.half[1] == b.tail {
		domino.playedFlipped = true
		return domino.half[0]
	}
	return -1
}

// NewBoard - new instance
func NewBoard(first *Domino) (*Board, error) {
	if first == nil {
		return nil, errors.New("Not a valid domino")
	}
	var board Board
	board.played = []*Domino{first}

	board.head = first.half[0]
	board.tail = first.half[1]

	return &board, nil
}
