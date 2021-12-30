package tm_test

import (
	"testing"

	. "github.com/arunmurugan78/tm"
	tp "github.com/arunmurugan78/tm/tape"
	"github.com/stretchr/testify/assert"
)

func TestUsingToUppercaseProgram(t *testing.T) {
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

func TestUsingEqualAandBProgram(t *testing.T) {
	transitions := make(TransitionMap)

	transitions["Q0"] = []Transition{
		{Direction: RightDirection, ToState: "Q0", ReadSymbol: byte('x'), WriteSymbol: byte('x')},
		{Direction: RightDirection, ToState: "Q0", ReadSymbol: byte('y'), WriteSymbol: byte('y')},
		{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('a'), WriteSymbol: byte('x')},
		{Direction: RightDirection, ToState: "Q3", ReadSymbol: byte('b'), WriteSymbol: byte('y')},
	}

	transitions["Q1"] = []Transition{
		{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('x'), WriteSymbol: byte('x')},
		{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('y'), WriteSymbol: byte('y')},
		{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('a'), WriteSymbol: byte('a')},
		{Direction: LeftDirection, ToState: "Q2", ReadSymbol: byte('b'), WriteSymbol: byte('y')},
	}

	transitions["Q3"] = []Transition{
		{Direction: RightDirection, ToState: "Q3", ReadSymbol: byte('x'), WriteSymbol: byte('x')},
		{Direction: RightDirection, ToState: "Q3", ReadSymbol: byte('y'), WriteSymbol: byte('y')},
		{Direction: RightDirection, ToState: "Q3", ReadSymbol: byte('b'), WriteSymbol: byte('b')},
		{Direction: LeftDirection, ToState: "Q2", ReadSymbol: byte('a'), WriteSymbol: byte('x')},
	}

	transitions["Q2"] = []Transition{
		{Direction: LeftDirection, ToState: "Q2", ReadSymbol: byte('x'), WriteSymbol: byte('x')},
		{Direction: LeftDirection, ToState: "Q2", ReadSymbol: byte('y'), WriteSymbol: byte('y')},
		{Direction: LeftDirection, ToState: "Q2", ReadSymbol: byte('a'), WriteSymbol: byte('a')},
		{Direction: LeftDirection, ToState: "Q2", ReadSymbol: byte('b'), WriteSymbol: byte('b')},
		{Direction: RightDirection, ToState: "Q0", ReadSymbol: byte('$'), WriteSymbol: byte('$')},
	}

	tm1 := NewTM(Config{
		StartState:    "Q0",
		AcceptedState: "Q0",
		InputString:   "aaabaabbbb",
		Transitions:   transitions,
	})

	tm1.Run()

	assert.Equal(t, tm1.IsAccepted(), true)

	tm2 := NewTM(Config{
		StartState:    "Q0",
		AcceptedState: "Q0",
		InputString:   "abaabbb",
		Transitions:   transitions,
	})

	tm2.Run()

	assert.Equal(t, tm2.IsAccepted(), false)

}
