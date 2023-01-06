package fizzbuzz

import "strconv"

func Convert(n int) string {
	if n % 3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n);
}