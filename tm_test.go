package tm_test

import (
	"testing"

	. "github.com/arunmurugan78/tm"
	tp "github.com/arunmurugan78/tm/tape"
	"github.com/stretchr/testify/assert"
)

func TestTM(t *testing.T) {
	transitions := make(TransitionMap)

	transitions["Q1"] = []Transition{{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('a'), WriteSymbol: byte('A')},
		{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('b'), WriteSymbol: byte('B')}}

	tm := NewTM(Config{
		StartState:    "Q1",
		AcceptedState: "Q1",
		InputString:   "abaabab",
		Transitions:   transitions,
	})

	tm.Run()

	tape := tm.GetTape()

	str := ""

	tape.MoveLeft()

	for tape.ReadSymbol() != tp.BlankSymbol {
		str = string(tape.ReadSymbol()) + str
		tape.MoveLeft()
	}

	assert.Equal(t, str, "ABAABAB")
}
