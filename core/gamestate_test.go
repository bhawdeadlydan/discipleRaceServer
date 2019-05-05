package core

import "testing"

func TestGameState_getAsJson(t *testing.T) {
	type fields struct {
		PlayerPositions  map[string]*Position
		PlayerAttributes map[string]*Attributes
	}
	positionOne := &Position{
		X:1,
		Y:1,
	}
	positionTwo := &Position{
		X:2,
		Y:2,
	}

	playerOneAttributes := &Attributes{
		Name: "1",
	}
	playerTwoAttributes := &Attributes{
		Name: "2",
	}

	playerPositions := make(map[string]*Position)
	playerPositions["1"] = positionOne
	playerPositions["2"] = positionTwo


	playerAttributes := make(map[string]*Attributes)
	playerAttributes["1"]= playerOneAttributes
	playerAttributes["2"]= playerTwoAttributes
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	 {
	 	name:"test1",
	 	wantErr:false,
	 	want: `{"PlayerPositions":{"1":{"X":1,"Y":1},"2":{"X":2,"Y":2}},"PlayerAttributes":{"1":{"Name":"1"},"2":{"Name":"2"}}}`,
	 	fields: fields{
	 		PlayerPositions:playerPositions,
	 		PlayerAttributes:playerAttributes,
		},
	 },
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GameState{
				PlayerPositions:  tt.fields.PlayerPositions,
				PlayerAttributes: tt.fields.PlayerAttributes,
			}
			got, err := gs.getAsJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("GameState.getAsJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GameState.getAsJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
