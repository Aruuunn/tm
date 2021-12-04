package tape_test

import (
	"testing"

	. "github.com/arunmurugan78/tm/tape"
	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/assert"
)

// makes sure that the RTape interface is relevant to Tape struct
func TestRTapeImplementedByTape(t *testing.T) {
	var rtape RTape = NewTape()
	rtape.MoveRight()
}

func TestInfinteTape(t *testing.T) {
	tape := NewTape()

	for i := 0; i < 5000; i++ {
		tape.MoveRight()
		assert.Equal(t, tape.ReadSymbol(), BlankSymbol)
	}

	for i := 0; i < 15000; i++ {
		tape.MoveLeft()
		assert.Equal(t, tape.ReadSymbol(), BlankSymbol)
	}
}

func TestTapeRightWriting(t *testing.T) {
	arr := make([]string, 5000)

	for idx := range arr {
		arr[idx] = uniuri.New()
	}

	tape := NewTape()

	for i := 0; i < 5000; i++ {
		tape.WriteSymbol(arr[i][0])
		tape.MoveRight()
	}

	tape.MoveLeft()

	for i := 5000 - 1; i >= 0; i-- {
		assert.Equal(t, arr[i][0], tape.ReadSymbol())
		tape.MoveLeft()
	}
}

func TestTapeLeftWriting(t *testing.T) {
	arr := make([]string, 5000)

	for idx := range arr {
		arr[idx] = uniuri.New()
	}

	tape := NewTape()

	for i := 0; i < 5000; i++ {
		tape.WriteSymbol(arr[i][0])
		tape.MoveLeft()
	}

	tape.MoveRight()

	for i := 5000 - 1; i >= 0; i-- {
		assert.Equal(t, arr[i][0], tape.ReadSymbol())
		tape.MoveRight()
	}
}

func TestLeftAndRightWriting(t *testing.T) {
	arr := make([]string, 15000)

	for idx := range arr {
		arr[idx] = uniuri.New()
	}

	tape := NewTape()

	for i := 0; i < 5000; i++ {
		tape.WriteSymbol(arr[i+10000][0])
		tape.MoveRight()
	}

	for i := 0; i < 5000; i++ {
		tape.MoveLeft()
	}

	tape.MoveLeft()

	for i := 10000 - 1; i >= 0; i-- {
		tape.WriteSymbol(arr[i][0])
		tape.MoveLeft()
	}

	tape.MoveRight()

	for i := 0; i < 15000; i++ {
		assert.Equal(t, tape.ReadSymbol(), arr[i][0])
		tape.MoveRight()
	}
}
