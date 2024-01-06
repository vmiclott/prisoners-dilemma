package main

import (
	"fmt"
	"math/rand"
	"prisoners-dilemma/game"
	"prisoners-dilemma/strategies"
)

func main() {
	numRounds := 200
	playersMap := map[string]*game.Player{
		"AlwaysCooperate":      game.NewPlayer(strategies.AlwaysCooperate{}),
		"Grofman":              game.NewPlayer(strategies.Grofman{Rand: rand.New(rand.NewSource(1))}),
		"InterchangeCooperate": game.NewPlayer(strategies.InterchangeCooperate{}),
		"Joss":                 game.NewPlayer(strategies.Joss{Rand: rand.New(rand.NewSource(2))}),
		"Random":               game.NewPlayer(strategies.Random{Rand: rand.New(rand.NewSource(3))}),
		"TitForTat":            game.NewPlayer(strategies.TitForTat{}),
		"AlwaysDefect":         game.NewPlayer(strategies.AlwaysDefect{}),
		"Grudger":              game.NewPlayer(strategies.Grudger{}),
		"InterchangeDefect":    game.NewPlayer(strategies.InterchangeDefect{}),
		"Pavlov":               game.NewPlayer(strategies.Pavlov{}),
		"SuspiciousTitForTat":  game.NewPlayer(strategies.SuspiciousTitForTat{}),
		"TitForTwoTats":        game.NewPlayer(strategies.TitForTwoTats{}),
	}
	players := []*game.Player{}
	for _, player := range playersMap {
		players = append(players, player)
	}
	rules := game.Rules{
		BothDefectGain:    1,
		BothCooperateGain: 3,
		WinGain:           5,
		LossGain:          0,
	}
	tournament := game.NewTournament(players, numRounds, rules)
	tournament.Simulate()
	for name, player := range playersMap {
		fmt.Printf("Player %s score: %d\n", name, player.Score)
	}
}
