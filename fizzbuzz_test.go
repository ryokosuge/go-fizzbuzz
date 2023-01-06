package fizzbuzz_test

import (
	"testing"

	"github.com/ryokosuge/go-fizzbuzz"
)

func TestConvert(t *testing.T) {

	tests := []struct{
		n int
		want string
	}{
		{
			n:    1,
			want: "1",
		},
		{
			n:    2,
			want: "2",
		},
		{
			n:    3,
			want: "Fizz",
		},
		{
			n:    5,
			want: "Buzz",
		},
		{
			n:    15,
			want: "FizzBuzz",
		},
	}

	for _, tt := range tests {
		got := fizzbuzz.Convert(tt.n)
		if got != tt.want {
			t.Errorf("Convert(%v) = %q but want %q", tt.n, got, tt.want)
		}
	}
}
