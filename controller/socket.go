package controller

import (
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/discipleRaceServer/logger"
	"github.com/discipleRaceServer/utils"
	"github.com/discipleRaceServer/model"
	"github.com/discipleRaceServer/core"
)

// BroadcastRequest struct
type BroadcastRequest struct {
	Channels []string `json:"channels"`
	Data     string   `json:"data"`
}

// PublishRequest struct
type PublishRequest struct {
	Channel string `json:"channel"`
	Data    string `json:"data"`
}

// Websocket Object
type Websocket struct {
	Clients      map[string]*websocket.Conn
	GameHandler  *core.GameHandler
	MatchRequest chan core.MatchRequest
	Broadcast    chan model.GameAction
	Upgrader     websocket.Upgrader
}

func NewWebsocket(gameHandler *core.GameHandler, matchRequest chan core.MatchRequest) *Websocket {
	return &Websocket{
		GameHandler:  gameHandler,
		MatchRequest: matchRequest,
	}
}

// Init initialize the websocket object
func (wsHandler *Websocket) Init() {
	wsHandler.Clients = make(map[string]*websocket.Conn)
	wsHandler.Broadcast = make(chan model.GameAction, 100)
	wsHandler.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}
}

// HandleConnections manage new clients
func (wsHandler *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request, playerID string, token string, correlationID string) {
	validate := utils.Validator{}
	if validate.IsEmpty(playerID) || validate.IsEmpty(token) || !validate.IsUUID4(playerID) {
		return
	}

	// Upgrade initial GET request to a websocket
	ws, err := wsHandler.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		logger.Fatalf(
			`Error while upgrading the GET request to a websocket for client %s: %s {"correlationId":"%s"}`,
			playerID,
			err.Error(),
			correlationID,
		)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	wsHandler.Clients[playerID] = ws

	logger.Infof(
		`Client %s connected {"correlationId":"%s"}`,
		playerID,
		correlationID,
	)

	cn, ok := w.(http.CloseNotifier)
	if !ok {
		http.NotFound(w, r)
		return
	}

	for {
		var gameAction model.GameAction

		err := ws.ReadJSON(&gameAction)

		if err != nil {
			delete(wsHandler.Clients, playerID)
			logger.Infof(
				`Client %s disconnected {"correlationId":"%s"}`,
				playerID,
				correlationID,
			)
			break
		}

		matchRequestType := ""
		if gameAction.Type == model.ActionType_Lifecycle && gameAction.Data == "join" {
			matchRequestType = core.MatchRequestType_Join
		}

		matchRequest := core.MatchRequest{
			PlayerID:         playerID,
			PlayerConnection: ws,
			Type:             matchRequestType,
		}
		wsHandler.MatchRequest <- matchRequest
		select {
		case <-cn.CloseNotify():
			logger.Infof("Client disconnected")
			delete(wsHandler.Clients, playerID)
			break
		}
	}
	logger.Infof("end of websocket connection!!")
}
