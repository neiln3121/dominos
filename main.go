package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/neiln3121/dominos/delivery"
	"github.com/neiln3121/dominos/game"
)

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		exitGame()
		os.Exit(0)
	}()

	noOfPlayers := flag.Int("players", 2, "number of players")
	flag.Parse()
	if *noOfPlayers > 4 {
		log.Fatal("Cannot have more than 4 players")
	}

	dominoGame := game.NewGame(*noOfPlayers)
	fmt.Println("Preparing Game")
	err := delivery.SelectPlayerNames(dominoGame)
	if err != nil {
		log.Print("Error setting up player names")
	}
	for !dominoGame.IsSetup() {
		delivery.SelectInitialDominos(dominoGame)
	}

	fmt.Println("Starting Game")
	err = dominoGame.StartGame()
	if err != nil {
		log.Fatalf("Could not start the game: %v", err)
	}

	for !dominoGame.IsFinished() {
		delivery.SelectOption(dominoGame)
	}
	winner, winnerTotal := dominoGame.GetWinner()
	fmt.Printf("WINNER: Player %s with %d score\n\n", winner, winnerTotal)
}

func exitGame() {
	fmt.Println("\n\nNo Winner !")
}
