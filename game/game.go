package game

type Rules struct {
	BothDefectGain    int
	BothCooperateGain int
	WinGain           int
	LossGain          int
}

type Game struct {
	player1   *Player
	player2   *Player
	numRounds int
	rules     Rules
}

func NewGame(player1 *Player, player2 *Player, numRounds int, rules Rules) *Game {
	return &Game{
		player1:   player1,
		player2:   player2,
		numRounds: numRounds,
		rules:     rules,
	}
}

func (game *Game) Play() {
	for i := 0; i < game.numRounds; i++ {
		player1Choices := game.player1.choices
		player2Choices := game.player2.choices
		player1Choice := game.player1.Choose(player2Choices)
		player2Choice := game.player2.Choose(player1Choices)
		results := getRoundResults(player1Choice, player2Choice, game.rules)
		game.player1.AddToScore(results[0])
		game.player2.AddToScore(results[1])
	}
}

func getRoundResults(player1Choice Choice, player2Choice Choice, rules Rules) [2]int {
	if player1Choice == player2Choice {
		if player1Choice == Defect {
			return [2]int{rules.BothDefectGain, rules.BothDefectGain}
		} else {
			return [2]int{rules.BothCooperateGain, rules.BothCooperateGain}
		}
	}
	if player1Choice == Defect {
		return [2]int{rules.WinGain, rules.LossGain}
	} else {
		return [2]int{rules.LossGain, rules.WinGain}
	}
}
