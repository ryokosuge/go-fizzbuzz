package fizzbuzz_test

import (
	"testing"

	"github.com/ryokosuge/go-fizzbuzz"
)

func TestFizz(t *testing.T) {
	got := fizzbuzz.Convert(1)
	if got != "1" {
		t.Errorf("Convert(1) is %q", got)
	}
}
