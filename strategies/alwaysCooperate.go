package strategies

import "prisoners-dilemma/game"

type AlwaysCooperate struct {
}

func (alwaysCooperate AlwaysCooperate) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	return game.Cooperate
}
