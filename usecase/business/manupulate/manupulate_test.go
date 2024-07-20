package manupulate

import (
	"fmt"
	"testing"
)

// Write your test here

func TestManupulate(t *testing.T) {
	//data := []string{"a"}
	//data := []string{"ab"}
	data := []string{"abc"}
	//data := []string{"aabb"}

	var result []string
	result = Manipulate(data)
	fmt.Println(result)

}
