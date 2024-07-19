package main

import (
	"fmt"
	"testing"
)

// Write your test here

func TestDDD(t *testing.T) {
	//data := []string{"a"}
	//data := []string{"ab"}
	data := []string{"aabb"}
	//data := []string{"aabb"}

	var result []string
	result = Manipulate(data)
	fmt.Println(result)
}
