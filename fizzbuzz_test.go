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

	got = fizzbuzz.Convert(2)
	if got != "2" {
		t.Errorf("Convert(2) is %q", got)
	}

	got = fizzbuzz.Convert(3)
	if got != "Fizz" {
		t.Errorf("Convert(3) is %q", got)
	}
}
