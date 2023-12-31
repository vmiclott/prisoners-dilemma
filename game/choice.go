package game

type Choice int

type Choices []Choice

const (
	Defect Choice = iota
	Cooperate
)
