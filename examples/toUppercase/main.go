package main

import (
	"fmt"

	. "github.com/arunmurugan78/tm"
	tp "github.com/arunmurugan78/tm/tape"
)

func main() {
	var s string

	fmt.Println("Enter a string consisting of only 'a' and 'b's :")
	fmt.Scanf("%s", &s)

	transitions := make(TransitionMap)

	transitions["Q1"] = []Transition{{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('a'), WriteSymbol: byte('A')},
		{Direction: RightDirection, ToState: "Q1", ReadSymbol: byte('b'), WriteSymbol: byte('B')}}

	tm := NewTM(Config{
		StartState:    "Q1",
		AcceptedState: "Q1",
		InputString:   s,
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

	fmt.Println("OUTPUT: ", str)
}
