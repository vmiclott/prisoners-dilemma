package strategies

import "prisoners-dilemma/game"

type SuspiciousTitForTat struct {
}

func (suspiciousTitForTat SuspiciousTitForTat) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(opponentsChoices)
	if n > 0 {
		return opponentsChoices[n-1]
	}
	return game.Defect
}
