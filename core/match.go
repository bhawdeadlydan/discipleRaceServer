package core

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type Match struct {
	ID                string
	PlayerConnections map[string]*websocket.Conn
	GameState         *GameState
}

func NewMatch(playerConnections map[string]*websocket.Conn) *Match {

	players := make([]string, 0, len(playerConnections))
	for playerID, _ := range playerConnections {
		players = append(players, playerID)
	}

	return &Match{
		ID:                uuid.NewV4().String(),
		PlayerConnections: playerConnections,
		GameState:         NewGameStateFor(players),
	}
}

// has to be a goroutine
func (m *Match) Start() {
	// start the match

	// updating the game state in constant interval

	// updating game state based on incoming actions from the players

	// broadcast the game state after every update to the players
}
