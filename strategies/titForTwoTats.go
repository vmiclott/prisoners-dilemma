package strategies

import "prisoners-dilemma/game"

type TitForTwoTats struct {
}

func (titForTwoTats TitForTwoTats) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(opponentsChoices)
	if n > 1 && opponentsChoices[n-2] == opponentsChoices[n-1] && opponentsChoices[n-1] == game.Defect {
		return game.Defect
	}
	return game.Cooperate
}
