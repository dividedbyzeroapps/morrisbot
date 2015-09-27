package main

// http://rhymebrain.com/talk?function=getRhymes&word=coin&maxResults=0&lang=en

import (
	// "bufio"
	"fmt"
	// "io"
	"io/ioutil"
	// "os"
)

// Gross
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	phraseFiles := []string{
		"phrases/beatles_songs.txt",
		"phrases/best-selling-books.txt",
		"phrases/movie-quotes.txt",
		"phrases/oscar_winning_movies.txt",
		"phrases/wikipedia_idioms.txt",
	}
	phraseLines := []string{}

	for _, fileName := range phraseFiles {
		dat, err := ioutil.ReadFile(fileName)
		check(err)
		phraseLines = append(phraseLines, string(dat))
	}
	fmt.Println(phraseLines)
}
