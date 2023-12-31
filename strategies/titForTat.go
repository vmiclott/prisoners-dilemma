package strategies

import "prisoners-dilemma/game"

type TitForTat struct {
}

func (titForTat TitForTat) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(opponentsChoices)
	if n > 0 {
		return opponentsChoices[n-1]
	}
	return game.Cooperate
}
