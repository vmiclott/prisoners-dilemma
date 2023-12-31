package strategies

import "prisoners-dilemma/game"

type AlwaysDefect struct {
}

func (alwaysDefect AlwaysDefect) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	return game.Defect
}
