package game

type Strategy interface {
	Choose(choices Choices, opponentsChoices Choices) Choice
}
