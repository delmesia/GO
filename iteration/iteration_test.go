package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func PrefixTrim(s string) string {
	s = strings.TrimPrefix(s, "!!!")
	return s
}

func TestPrefixTrim(t *testing.T) {
	got := PrefixTrim("!!!123")
	want := "123"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Repeat(character string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected  %q got  %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 7)
	fmt.Println(repeated)
	//Output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
