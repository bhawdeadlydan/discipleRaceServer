package core

import (
	"encoding/json"
)

type Position struct {
	X int
	Y int
}

type Attributes struct {
	Name string
}

func NewPosition(X int, Y int) *Position {
	return &Position{
		X: X,
		Y: Y,
	}
}

func NewAttributes(playerID string) *Attributes {
	return &Attributes{
		Name: playerID,
	}
}

type GameState struct {
	PlayerPositions  map[string]*Position
	PlayerAttributes map[string]*Attributes
}

func (gs *GameState) getAsJson() (string, error) {
	state, _ := json.Marshal(gs)
	return string(state), nil
}

func NewGameStateFor(players []string) *GameState {
	gameState := &GameState{
		PlayerPositions:  make(map[string]*Position, 0),
		PlayerAttributes: make(map[string]*Attributes, 0),
	}

	for index, player := range players {
		gameState.PlayerPositions[player] = NewPosition(index, index)
		gameState.PlayerAttributes[player] = NewAttributes(player)
	}

	return gameState
}
