package main

import (
	"fmt"
	"prisoners-dilemma/game"
	"prisoners-dilemma/strategies"
)

func main() {
	numRounds := 200
	player1 := game.NewPlayer(strategies.TitForTat{})
	player2 := game.NewPlayer(strategies.AlwaysDefect{})
	rules := game.Rules{
		BothDefectGain:    1,
		BothCooperateGain: 3,
		WinGain:           5,
		LossGain:          0,
	}
	game := game.NewGame(player1, player2, numRounds, rules)
	game.Simulate()
	fmt.Printf("Player 1 score: %d\n", player1.Score)
	fmt.Printf("Player 2 score: %d\n", player2.Score)
}
