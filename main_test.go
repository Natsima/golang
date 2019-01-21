package main

import "testing"

func Test_1_should_be_1(t *testing.T) {
	result := Fizzbuzz("1")
	if result != "1" {
		t.Error("1 shoould be 1 but got", result)
	}
}
func Test_2_should_be_2(t *testing.T) {
	result := Fizzbuzz("2")
	if result != "2" {
		t.Error("2 shoould be 2 but got", result)
	}
}
func Test_3_should_be_Fizz(t *testing.T) {
	result := Fizzbuzz("3")
	if result != "Fizz" {
		t.Error("3 shoould be Fizz but got", result)
	}
}
func Test_4_should_be_4(t *testing.T) {
	result := Fizzbuzz("4")
	if result != "4" {
		t.Error("4 shoould be 4 but got", result)
	}
}

func Test_5_should_be_Buzz(t *testing.T) {
	result := Fizzbuzz("5")
	if result != "Buzz" {
		t.Error("5 shoould be Buzz but got", result)
	}
}

func Test_6_should_be_Fizz(t *testing.T) {
	result := Fizzbuzz("6")
	if result != "Fizz" {
		t.Error("6 shoould be Fizz but got", result)
	}
}
