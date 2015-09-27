package main

// http://rhymebrain.com/talk?function=getRhymes&word=coin&maxResults=0&lang=en

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Gross
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type RhymebrainResult struct {
	Word      string `json:"word"`
	Freq      int    `json:"freq"`
	Score     int    `json:"score"`
	Flags     string `json:"flags"`
	Syllables string `json:"syllables"`
}

// Rhymes are in rhyme_service.rb

func rhymes(word []byte) {
	many := []byte(`
		[ {"word":"goin","freq":19,"score":300,"flags":"c","syllables":"1"},
		{"word":"join","freq":23,"score":300,"flags":"bc","syllables":"1"},
		{"word":"groin","freq":18,"score":300,"flags":"bc","syllables":"1"},
		{"word":"loin","freq":18,"score":300,"flags":"bc","syllables":"1"},
		{"word":"doyenne","freq":13,"score":300,"flags":"b","syllables":"1"},
		{"word":"groyne","freq":13,"score":300,"flags":"b","syllables":"1"},
		{"word":"quoin","freq":13,"score":300,"flags":"b","syllables":"1"},
		{"word":"adjoin","freq":16,"score":300,"flags":"bc","syllables":"2"},
		{"word":"enjoin","freq":18,"score":300,"flags":"bc","syllables":"2"},
		{"word":"rejoin","freq":18,"score":300,"flags":"bc","syllables":"2"},
		{"word":"conjoin","freq":15,"score":300,"flags":"bc","syllables":"2"},
		{"word":"purloin","freq":14,"score":300,"flags":"bc","syllables":"2"},
		{"word":"subjoin","freq":16,"score":300,"flags":"b","syllables":"2"},
		{"word":"sirloin","freq":15,"score":300,"flags":"b","syllables":"2"},
		{"word":"tenderloin","freq":15,"score":300,"flags":"bc","syllables":"3"}]
	`)
	results := make([]RhymebrainResult, 0)
	err := json.Unmarshal(many, &results)

	check(err)
	fmt.Println(results)
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
	rhymes([]byte("coin"))
	// fmt.Println(phraseLines)
}
