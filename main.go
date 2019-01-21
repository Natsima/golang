package main

import "strconv"

func Fizzbuzz(number string) string {
	x, _ := strconv.Atoi(number)

	if x%3 == 0 && x%5 == 0 {
		return "Fizzbuzz"
	}
	if x%3 == 0{
		return "Fizz"
	}

	if x%5 == 0 {
		return "Buzz"
	}

	return number
}
