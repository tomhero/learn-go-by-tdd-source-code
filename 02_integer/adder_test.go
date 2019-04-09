package integer

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestAdderWithNegative(t *testing.T) {
	sum := Add(-2, -4)
	expected := -6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
