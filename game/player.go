package game

type Player struct {
	choices  Choices
	strategy Strategy
	Score    int
}

func NewPlayer(strategy Strategy) *Player {
	return &Player{
		choices:  Choices{},
		strategy: strategy,
	}
}

func (player *Player) Choose(opponentsChoices Choices) Choice {
	choice := player.strategy.Choose(player.choices, opponentsChoices)
	player.choices = append(player.choices, choice)
	return choice
}

func (player *Player) AddToScore(score int) {
	player.Score += score
}
