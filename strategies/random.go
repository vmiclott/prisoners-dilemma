package strategies

import (
	"math/rand"
	"prisoners-dilemma/game"
)

type Random struct {
	Rand *rand.Rand
}

func (random Random) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	if random.Rand.Intn(2) < 1 {
		return game.Defect
	}
	return game.Cooperate
}
