# Prisoner's dilemma

Based on https://en.wikipedia.org/wiki/Prisoner%27s_dilemma.

## Game configuration

### Create players

Player strategies implement the following interface:

```go
type Strategy interface {
	Choose(choices Choices, opponentsChoices Choices) Choice
}
```

Some common strategies such as always defect and tit-for-tat are already implemented.

Players can be created using the `NewPlayer` function which requires a strategy.

```go
func NewPlayer(strategy Strategy) *Player
```

### Create rule set

The game rule set defines scores for all possible game round outcomes

```go
type Rules struct {
	BothDefectGain    int
	BothCooperateGain int
	WinGain           int
	LossGain          int
}
```

### Create Game

A game can be created using the `NewGame` function which requires 2 players, a number of rounds and a ruleset.

```go
func NewGame(player1 *Player, player2 *Player, numRounds int, rules Rules) *Game
```

### Create Tournament

A tournament can be created using the `NewTournament` function which requires a slice of players, a number of rounds and a ruleset. A tournament is a set of games with exactly one game for each pair of players.

```go
func NewGame(players []*Player, numRounds int, rules Rules) *Tournament
```

## Game simulation

After a game is created, use `Simulate` to run a simulation.

```go
func (game *Game) Simulate()
```

Each player's score can be accessed using the public `Score` field.

## Tournament simulation

After a tournament is created, use `Simulate` to run a simulation.

```go
func (tournament *Tournament) Simulate()
```

Each player's score can be accessed using the public `Score` field.
