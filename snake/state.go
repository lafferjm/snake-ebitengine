package snake

type State int

const (
	MainMenu State = iota
	Playing
	GameOver
)

type GameState struct {
	state State
}

func NewGameState() *GameState {
	return &GameState{
		state: MainMenu,
	}
}

func (gs *GameState) SetState(state State) {
	gs.state = state
}

func (gs *GameState) PollState(state State) bool {
	return gs.state == state
}
