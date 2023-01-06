package fizzbuzz

import "strconv"

func Convert(n int) string {
	if n % 5 == 0 && n % 3 == 0 {
		return "FizzBuzz"
	}

	if n % 5 == 0 {
		return "Buzz"
	}

	if n % 3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(n);
}