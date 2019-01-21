package main

import "strconv"

func Fizzbuzz(number string) string {
	x, _ := strconv.Atoi(number)

	if x%3 == 0 {
		return "Fizz"
	}

	if number == "5" {
		return "Buzz"
	}

	return number
}
