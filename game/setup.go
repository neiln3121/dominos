package game

import (
	"math/rand"
	"time"

	"github.com/neiln3121/dominos/models"
)

func initPlayers(noOfPlayers int) []*models.Player {
	players := make([]*models.Player, noOfPlayers)
	for i := 0; i < noOfPlayers; i++ {
		players[i] = models.NewPlayer()
	}
	return players
}

func initTable() *models.Table {
	table := models.NewTable()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(table.Dominos), func(i, j int) { table.Dominos[i], table.Dominos[j] = table.Dominos[j], table.Dominos[i] })

	return table
}

func initBoard(players []*models.Player) (board *models.Board, firstPlayerIndex int, err error) {
	firstPlayerIndex = -1
	firstPlayerDominoIndex := -1
	highest := -1

	for i, player := range players {
		playerHighest, dominoIndex := player.GetHighestDouble()
		if playerHighest > highest {
			highest = playerHighest
			firstPlayerIndex = i
			firstPlayerDominoIndex = dominoIndex
		}
	}
	// If no doubles then play a random domino from a random player
	if highest == -1 {
		firstPlayerIndex = rand.Intn(len(players) - 1)
		firstPlayerDominoIndex = rand.Intn(players[firstPlayerIndex].DominoCount() - 1)
	}
	highestDomino, err := players[firstPlayerIndex].Get(firstPlayerDominoIndex)
	if err != nil {
		return
	}
	players[firstPlayerIndex].RemoveDomino(firstPlayerDominoIndex)
	board, err = models.NewBoard(highestDomino)
	return
}
