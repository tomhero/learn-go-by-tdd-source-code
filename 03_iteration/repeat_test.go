package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("T", 5)
	expected := "TTTTT"

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func TestAnotherRepeat(t *testing.T) {
	repeated := Repeat("R", 5)
	expected := "RRRRR"

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func TestRepeatWithOnlyOneTime(t *testing.T) {
	repeated := Repeat("S", 1)
	expected := "S"

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func TestRepeatWithZeroCount(t *testing.T) {
	repeated := Repeat("z", 0)
	expected := ""

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func TestRepeatWithNegativeValue(t *testing.T) {
	repeated := Repeat("N", -5)
	expected := ""

	if repeated != expected {
		t.Errorf("expected '%s' got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
