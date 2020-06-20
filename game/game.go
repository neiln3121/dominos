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

// SetNextPlayer - sets the next player to play a round
func (g *Game) SetNextPlayer() {
	g.currentPlayerIndex++
	if g.currentPlayerIndex >= len(g.players) {
		g.currentPlayerIndex = 0
	}
}

// IsFinished - checks whether anyone has one or if everyone can proceed
func (g *Game) IsFinished() bool {
	if g.GetCurrentPlayer().DominoCount() == 0 {
		g.winner = g.GetCurrentPlayer()
		g.status = finished
		return true
	}

	if g.table.AllPicked() {
		canProceed := false
		playersCanProceed := 0
		for _, player := range g.players {
			for _, domino := range player.GetDominos() {
				if domino.Half[0] == g.board.GetHead() || domino.Half[0] == g.board.GetTail() ||
					domino.Half[1] == g.board.GetHead() || domino.Half[1] == g.board.GetTail() {
					canProceed = true
					break
				}
			}
			if canProceed {
				playersCanProceed++
			}
		}

		if playersCanProceed != len(g.players) {
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
func (g *Game) GetWinner() (playerID int, playerTotal int) {
	if g.status == finished {
		playerID = g.winner.ID
		playerTotal = g.winner.TotalDots()
	}
	return
}

// PickDomino - pick a domino for a player from the table
func (g *Game) PickDomino(playerIndex, dominoID int) error {
	index := dominoID - 1
	domino, err := g.table.GetUnpickedDomino(index)
	if err != nil {
		return err
	}
	return g.players[playerIndex].AddDomino(domino)
}

// PlayDomino - play a domino for a player on the board
func (g *Game) PlayDomino(playerIndex, dominoID int, atHead bool) error {
	index := dominoID - 1
	domino, err := g.players[playerIndex].Get(index)
	if err != nil {
		return err
	}
	if err := g.board.PlayDomino(domino, atHead); err != nil {
		return err
	}

	return g.players[playerIndex].RemoveDomino(index)
}

// GetCurrentPlayer - get the player who is currently playing a round
func (g *Game) GetCurrentPlayer() *models.Player {
	return g.players[g.currentPlayerIndex]
}

// GetPlayers - get all players in the game
func (g *Game) GetPlayers() []*models.Player {
	return g.players
}

// GetNumberOfPlayers - get the number of players in the game
func (g *Game) GetNumberOfPlayers() int {
	return len(g.players)
}

// GetTable - get the table to pick dominos from
func (g *Game) GetTable() *models.Table {
	return g.table
}

// GetBoard - get the board to play dominos onto
func (g *Game) GetBoard() *models.Board {
	return g.board
}

// NewGame - new instance
func NewGame(noOfPlayers int) *Game {
	return &Game{
		players: initPlayers(noOfPlayers),
		table:   initTable(),
		status:  setup,
	}
}
