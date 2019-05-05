package model


const (
	ActionType_Control   = "control"
	ActionType_Lifecycle = "lifecycle"
)

// GameAction struct
type GameAction struct {
	Type string `json:"type"`
	Data string `json:"data"`

	FromClient string `json:"from_client"`
	ToClient   string `json:"to_client"`
	Channel    string `json:"channel"`
}

// IsValid checks if message is valid
func (m *GameAction) IsValid() bool {
	//validator := utils.Validator{}
	//return validator.IsJSON(m.Data)
	return true
}
