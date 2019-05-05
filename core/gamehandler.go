package core

type GameHandler struct {
	//action chan model.GameAction
	PlayerMatches map[string]*Match
	MatchRequest  chan MatchRequest
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		//action: make(chan model.GameAction, 0),
		PlayerMatches: make(map[string]*Match, 0),
		MatchRequest: make(chan MatchRequest, 100),
	}
}
