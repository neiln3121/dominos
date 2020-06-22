package game

import (
	"errors"

	"github.com/neiln3121/dominos/models"
)

type status int

const (
	setup status = iota
	inPlay
	finished
)

// Game - defines the rules and stores all the components for a game
type Game struct {
	players            []*models.Player
	table              *models.Table
	board              *models.Board
	currentPlayerIndex int
	winner             *models.Player
	status             status
}

// StartGame - start up a new game once setup
func (g *Game) StartGame() error {
	if g.status != setup {
		return errors.New("Game already in progess or complete")
	}
	board, firstPlayer, err := initBoard(g.players)
	if err != nil {
		return err
	}
	g.board = board
	g.currentPlayerIndex = firstPlayer
	g.status = inPlay

	return nil
}

// IsSetup - checks whether each player has enough dominos
func (g *Game) IsSetup() bool {
	g.setNextPlayer()

	return g.GetCurrentPlayer().HasStartingDominos()
}

// IsFinished - checks whether anyone has won or if everyone is able to proceed
func (g *Game) IsFinished() bool {
	if g.GetCurrentPlayer().DominoCount() == 0 {
		g.winner = g.GetCurrentPlayer()
		g.status = finished
		return true
	}

	g.setNextPlayer()
	if g.table.AllPicked() {
		// if current player can't proceed, calculate winner and end game
		if !g.GetCurrentPlayer().CanProceed(g.board.GetHead(), g.board.GetTail()) {
			minDots := 48
			for _, player := range g.players {
				playerDots := player.TotalDots()
				if playerDots < minDots {
					minDots = playerDots
					g.winner = player
				}
			}
			g.status = finished
		}
	}

	return g.status == finished
}

// GetWinner - get winner details
func (g *Game) GetWinner() (playerName string, playerTotal int) {
	if g.status == finished {
		playerName = g.winner.Name
		playerTotal = g.winner.TotalDots()
	}
	return
}

// PickDomino - pick a domino for a player from the table
func (g *Game) PickDomino(dominoID int) error {
	index := dominoID - 1
	domino, err := g.table.GetUnpickedDomino(index)
	if err != nil {
		return err
	}
	return g.players[g.currentPlayerIndex].AddDomino(domino)
}

// PlayDomino - play a domino for a player on the board
func (g *Game) PlayDomino(dominoID int, atHead bool) error {
	index := dominoID - 1
	domino, err := g.players[g.currentPlayerIndex].Get(index)
	if err != nil {
		return err
	}
	if err := g.board.PlayDomino(domino, atHead); err != nil {
		return err
	}

	return g.players[g.currentPlayerIndex].RemoveDomino(index)
}

// GetCurrentPlayer - get the player who is currently playing a round
func (g *Game) GetCurrentPlayer() *models.Player {
	return g.players[g.currentPlayerIndex]
}

// GetPlayers - get all players in the game
func (g *Game) GetPlayers() []*models.Player {
	return g.players
}

// GetTable - get the table to pick dominos from
func (g *Game) GetTable() *models.Table {
	return g.table
}

// GetBoard - get the board to play dominos onto
func (g *Game) GetBoard() *models.Board {
	return g.board
}

// SetNextPlayer - sets the next player to play a round
func (g *Game) setNextPlayer() {
	g.currentPlayerIndex++
	if g.currentPlayerIndex >= len(g.players) {
		g.currentPlayerIndex = 0
	}
}

// NewGame - new instance
func NewGame(noOfPlayers int) *Game {
	return &Game{
		players:            initPlayers(noOfPlayers),
		table:              initTable(),
		status:             setup,
		currentPlayerIndex: noOfPlayers - 1,
	}
}
