package strategies

import "prisoners-dilemma/game"

type Grudger struct {
	opponentHasDefected bool
}

func (grudger Grudger) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(opponentsChoices)
	if n > 1 && opponentsChoices[n-1] == game.Defect {
		grudger.opponentHasDefected = true
	}

	if grudger.opponentHasDefected {
		return game.Defect
	}

	return game.Cooperate
}
