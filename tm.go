package tm

import (
	"errors"

	. "github.com/arunmurugan78/tm/tape"
)

const (
	RightDirection = iota
	LeftDirection
)

type Transition struct {
	Direction   int
	ToState     string
	ReadSymbol  byte
	WriteSymbol byte
}

type TransitionMap map[string][]Transition

type TM struct {
	tape          *Tape
	currentState  string
	acceptedState string
	transitions   TransitionMap
}

type Config struct {
	Transitions   TransitionMap
	AcceptedState string
	StartState    string
	InputString   string
}

// WriteInputString writes the given string on the tape starting from the position pointed by the tape head.
func (tm *TM) WriteInputString(str string) {
	for i := 0; i < len(str); i++ {
		tm.tape.WriteSymbol(str[i])
		tm.tape.MoveRight()
	}

	for i := 0; i < len(str); i++ {
		tm.tape.MoveLeft()
	}
}

func (tm *TM) IsAccepted() bool {
	return tm.currentState == tm.acceptedState
}

func (tm *TM) GetTape() RTape {
	return tm.tape
}

func (tm *TM) GetCurrentState() string {
	return tm.currentState
}

func (tm *TM) getTransitionForCurrentState(readSymbol byte) (Transition, error) {
	transitions := tm.transitions[tm.currentState]

	for idx := range transitions {
		if transitions[idx].ReadSymbol == readSymbol {
			return transitions[idx], nil
		}
	}

	return Transition{}, errors.New("no transition available for the current state and the read symbol")
}

func (tm *TM) Run() {
	for tm.currentState != tm.acceptedState {
		symbol := tm.tape.ReadSymbol()
		transition, err := tm.getTransitionForCurrentState(symbol)

		if err != nil {
			// TM has entered the reject state
			return
		}

		tm.tape.WriteSymbol(transition.WriteSymbol)

		if transition.Direction == RightDirection {
			tm.tape.MoveRight()
		} else {
			tm.tape.MoveLeft()
		}

		tm.currentState = transition.ToState
	}
}

func NewTM(config Config) *TM {
	tm := &TM{
		tape:          NewTape(),
		acceptedState: config.AcceptedState,
		currentState:  config.StartState,
		transitions:   config.Transitions,
	}

	if config.InputString != "" {
		tm.WriteInputString(config.InputString)
	}

	return tm
}
