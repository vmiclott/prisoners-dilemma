package strategies

import "prisoners-dilemma/game"

type InterchangeCooperate struct {
}

func (interchangeCooperate InterchangeCooperate) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(choices)
	if n%2 == 0 {
		return game.Cooperate
	}
	return game.Defect
}
