package main

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
	Word  string `json:"word"`
	Score int    `json:"score"`
}

// Rhymes are in rhyme_service.rb

func rhymes(word []byte) {
	// TODO: get rhymes from
	// http://rhymebrain.com/talk?function=getRhymes&word=coin&maxResults=0&lang=en
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
	fmt.Println(results[0].Word)
}

func main() {
	keyword := "coin"
	// TODO:
	// 1) Iterate through all phrases and all rhymes,
	// 2) find phrases that have a rhyme in them
	// 3) replace the rhyme with the keyword in the phrase
	// 4) collect the changed phrases in a slice
	// 5) print out the changed phrases

	phraseFilePaths := []string{
		"phrases/beatles_songs.txt",
		"phrases/best-selling-books.txt",
		"phrases/movie-quotes.txt",
		"phrases/oscar_winning_movies.txt",
		"phrases/wikipedia_idioms.txt",
	}
	phrases := []string{}

	for _, path := range phraseFilePaths {
		contents, err := ioutil.ReadFile(path)
		check(err)
		phrases = append(phrases, string(contents))
	}
	rhymes([]byte(keyword))
	// fmt.Println(phraseLines)
}
