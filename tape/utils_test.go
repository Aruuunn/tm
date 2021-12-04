package tape_test

import (
	"testing"

	. "github.com/arunmurugan78/tm/tape"
)

func TestMax(t *testing.T) {
	if Max(0, 10) != 10 {
		t.Fatal("expected 10")
	}
}

func TestAbs(t *testing.T) {
	if Abs(-10) != 10 {
		t.Fatal("expected 10")
	}

	if Abs(100) != 100 {
		t.Fatal("expected 100")
	}

	if Abs(0) != 0 {
		t.Fatal("expected 0")
	}
}
