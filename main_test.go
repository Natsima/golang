package main

import "testing"

func Test_1_should_be_1(t *testing.T){
	result :=Fizzbuzz(1)
	if result != 1 { 
		t.Error("1 shoould be 1 but got",result)
	}	
}
