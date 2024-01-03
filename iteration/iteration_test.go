package iteration

import "testing"

func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated = repeated + character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected  %q got  %q", expected, repeated)
	}
}
