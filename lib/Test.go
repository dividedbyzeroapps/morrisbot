package main

import (
	"fmt"
	"regexp"
)

// Gross
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rhyme := "join"
	// re := regexp.MustCompile("(?i)\b" + rhyme + "\b")
	re := regexp.MustCompile("(?i)\\b" + rhyme + "\\b")
	matches := re.MatchString("Join up")
	var hasMatch string
	if matches {
		hasMatch = "true"
	} else {
		hasMatch = "false"
	}
	fmt.Println("Match? " + hasMatch)
}
