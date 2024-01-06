package game

type Tournament struct {
	players   []*Player
	numRounds int
	rules     Rules
}

func NewTournament(players []*Player, numRounds int, rules Rules) *Tournament {
	return &Tournament{
		players:   players,
		numRounds: numRounds,
		rules:     rules,
	}
}

func (tournament *Tournament) Simulate() {
	games := tournament.createGames()
	for _, game := range games {
		game.Simulate()
		tournament.resetPlayersChoices()
	}
}

func (tournament *Tournament) resetPlayersChoices() {
	for _, player := range tournament.players {
		player.ResetChoices()
	}
}

func (tournament *Tournament) createGames() []*Game {
	n := len(tournament.players)
	games := []*Game{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			games = append(games, NewGame(tournament.players[i], tournament.players[j], tournament.numRounds, tournament.rules))
		}
	}
	return games
}
