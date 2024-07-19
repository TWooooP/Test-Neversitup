package main

import (
	"strings"
)

func CountSmilyFace(text []string) int {
	// TODO : start your code here
	var CountResultData int
	for _, data := range text {
		//separate string
		sss := strings.Split(data, "")
		if ((sss[0] == ":" || sss[0] == ";") && (sss[1] == ")" || sss[1] == "D")) ||
			(sss[1] == "-" || sss[1] == "~" && (sss[1] == ")" || sss[1] == "D")) {
			CountResultData += 1
		}
	}
	return CountResultData

}
