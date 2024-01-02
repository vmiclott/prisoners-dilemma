package strategies

import (
	"math/rand"
	"prisoners-dilemma/game"
)

type Joss struct {
	Rand *rand.Rand
}

func (joss Joss) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(opponentsChoices)
	if n < 1 {
		return game.Cooperate
	}

	if opponentsChoices[n-1] == game.Cooperate && joss.Rand.Intn(10) < 9 {
		return game.Cooperate
	}

	return game.Defect
}
