package strategies

import "prisoners-dilemma/game"

type Pavlov struct {
}

func (pavlov Pavlov) Choose(choices game.Choices, opponentsChoices game.Choices) game.Choice {
	n, m := len(choices), len(opponentsChoices)

	if n > 1 {
		if choices[n-1] == opponentsChoices[m-1] {
			return game.Cooperate
		} else {
			return game.Defect
		}
	}

	return game.Cooperate
}
