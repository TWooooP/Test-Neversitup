package main

import (
	"fmt"
	"testing"
)

// Write your test here

func TestCheckOddNumber(t *testing.T) {
	//data := []int{7}
	//data := []int{1, 1, 2}
	//data := []int{0, 1, 0, 1, 0}
	//data := []int{0, 1, 0, 1, 0}
	data := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}

	result := FindOddNumber(data)

	fmt.Println(result)
}
