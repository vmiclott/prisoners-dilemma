package strategies

import "prisoners-dilemma/game"

type InterchangeDefect struct {
}

func (interchangeDefect InterchangeDefect) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n := len(choices)
	if n%2 == 0 {
		return game.Defect
	}
	return game.Cooperate
}
