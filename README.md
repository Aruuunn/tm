# tm

[![Go Reference](https://pkg.go.dev/badge/github.com/arunmurugan78/tm.svg)](https://pkg.go.dev/github.com/arunmurugan78/tm)
[![Go](https://github.com/ArunMurugan78/tm/actions/workflows/go.yml/badge.svg)](https://github.com/ArunMurugan78/tm/actions/workflows/go.yml)

A Golang package which implements [Turing Machine](https://en.wikipedia.org/wiki/Turing_machine).


## Installation
```
go get github.com/arunmurugan78/tm
```

## Quick Start

```go
// examples/equalAandB/main.go
package main

import (
	"fmt"

	. "github.com/arunmurugan78/tm"
)

func main() {
	var s string

	fmt.Println("Enter a string consisting of only 'a' and 'b's :")
	fmt.Scanf("%s", &s)

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

	tm := NewTM(Config{
		StartState:    "Q0",
		AcceptedState: "Q0",
		InputString:   s,
		Transitions:   transitions,
	})

	tm.Run()

	fmt.Println("has same no of 'a's and 'b's? ", tm.IsAccepted())
}

```

## License
This project is licensed under the MIT License.