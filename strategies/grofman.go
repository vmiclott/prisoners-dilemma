package strategies

import (
	"math/rand"
	"prisoners-dilemma/game"
)

type Grofman struct {
	Rand *rand.Rand
}

func (grofman Grofman) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(choices)
	if n < 2 {
		return game.Cooperate
	}
	if n < 7 {
		return opponentsChoices[n-1]
	}
	if choices[n-1] == opponentsChoices[n-1] {
		return game.Cooperate
	}

	if grofman.Rand.Intn(7) < 2 {
		return game.Cooperate
	}

	return game.Defect
}
