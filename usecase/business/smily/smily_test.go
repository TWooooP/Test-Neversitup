package smily

import (
	"fmt"
	"testing"
)

// Write your test here

func TestCheckSmile(t *testing.T) {
	//data := []string{":)", ";(", ";}", ":-D"}
	//data := []string{";D", ":-(", ":-)", ";~)"}
	data := []string{";]", ":[", ";*", ":$", ";-D"}
	result := CountSmilyFace(data)

	fmt.Println(result)
}
