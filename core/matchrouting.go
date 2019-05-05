package core

import (
	"time"
	"github.com/gorilla/websocket"
	"fmt"
	"github.com/tomdionysus/binarytree"
	"github.com/discipleRaceServer/helper"
)

const (
	MatchRequestType_Join   = "join"
	MatchRequestType_Create = "create"

	WALK_DIRECTION_BACKWARD = false
)

type MatchRequest struct {
	PlayerID         string
	PlayerConnection *websocket.Conn
	Type             string
}


type MatchRouting struct {
	PlayerStagingTree     *binarytree.Tree
	PlayerStagingRegister map[string]int64
	PlayerConnection      map[string]*websocket.Conn
	PlayerMatches         map[string]*Match
	MatchSizeMin          int
	// Match wait time in seconds is Skipped for now.
	// Implementing this would mean the remaining players have to be bots.
	MatchWaitTimeInSeconds int
	MatchSizMax            int
	MatchRequest           chan MatchRequest
}

func NewMatchRouting() *MatchRouting {
	playerStagingTree := binarytree.NewTree()
	return &MatchRouting{
		PlayerStagingTree:      playerStagingTree,
		PlayerStagingRegister:  make(map[string]int64, 0),
		PlayerConnection:       make(map[string]*websocket.Conn, 0),
		PlayerMatches:          make(map[string]*Match, 0),
		MatchSizeMin:           2,
		MatchWaitTimeInSeconds: 2,
		MatchRequest:           make(chan MatchRequest, 100),
		MatchSizMax:            5,
	}
}

func (mr *MatchRouting) GetMatchRequestChannel() chan MatchRequest {
	return mr.MatchRequest
}

func (mr *MatchRouting) AddPlayerToStagingArea(playerID string, playerConnection *websocket.Conn) {
	currentEpoch := time.Now().UnixNano()
	mr.PlayerStagingRegister[playerID] = currentEpoch
	mr.PlayerConnection[playerID] = playerConnection
	mr.PlayerStagingTree.Set(helper.Int64Key(currentEpoch), playerID)
}

// match staging area handles the players being added to a new match
// reconnecting players not handled for now
//TODO: handle reconnecting players after disconnection
func (mr *MatchRouting) Start() {
	for {
		matchRequest := <-mr.MatchRequest
		switch matchRequest.Type {
		case MatchRequestType_Join:
			mr.AddPlayerToStagingArea(matchRequest.PlayerID, matchRequest.PlayerConnection)
		case MatchRequestType_Create:
			mr.AssignMatchesToPlayers()
		default:
			fmt.Println(fmt.Sprintf("Unimplemented match request type !!- %#v", matchRequest))
		}

	}
}

func (mr *MatchRouting) AssignMatchesToPlayers() {
	batchSelector := NewBatchSelector(mr.MatchSizeMin, mr.MatchSizMax)
	treeIterator := func(key binarytree.Comparable, value interface{}) {
		currentPlayer := value.(string)
		batchSelector.AddPlayer(currentPlayer)
	}

	mr.PlayerStagingTree.Walk(treeIterator, WALK_DIRECTION_BACKWARD)
	selectedPlayersBatches := batchSelector.GetSelectedBatches()

	for _, selectedPlayers := range selectedPlayersBatches {
		selectedPlayersConnection := make(map[string]*websocket.Conn, 0)
		for _, player := range selectedPlayers {
			selectedPlayersConnection[player] = mr.PlayerConnection[player]
		}
		match := NewMatch(selectedPlayersConnection)
		match.Start()
		mr.assignedMatchTo(selectedPlayers, match)
	}

	for _, selectedPlayers := range selectedPlayersBatches {
		mr.removePlayersFromStaging(selectedPlayers)
	}
}

func (mr *MatchRouting) assignedMatchTo(selectedPlayers []string, assignedMatch *Match) {
	for _, player := range selectedPlayers {
		mr.PlayerMatches[player] = assignedMatch
		logMessage := fmt.Sprintf("Assigned Match %s to player [%s]", assignedMatch.ID, player)
		fmt.Println(logMessage)
	}
}

// delete those selectedPlayers from the waiting area.
func (mr *MatchRouting) removePlayersFromStaging(selectedPlayers []string) {
	for _, playerID := range selectedPlayers {
		playerRegistrationEpoch := mr.PlayerStagingRegister[playerID]
		mr.PlayerStagingTree.Clear(helper.Int64Key(playerRegistrationEpoch))
		delete(mr.PlayerStagingRegister, playerID)
		delete(mr.PlayerConnection, playerID)
	}
}


