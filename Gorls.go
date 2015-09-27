package main

// Print out puns for the provided word
// Usage:
//   go run Gorls.go heart
// or:
//   ./Gorls heart
//
// Inspired by/ripped off from https://github.com/iancanderson/girls_just_want_to_have_puns
// Thanks, Ian.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
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

type Pun struct {
	PunPhrase      string
	OriginalPhrase string
}

func rhymes(word string) []RhymebrainResult {
	url := "http://rhymebrain.com/talk?function=getRhymes&word=" + word + "&maxResults=0&lang=en"
	response, httpErr := http.Get(url)
	check(httpErr)
	defer response.Body.Close()
	rawJSON, readErr := ioutil.ReadAll(response.Body)
	check(readErr)
	results := make([]RhymebrainResult, 0)
	err := json.Unmarshal(rawJSON, &results)
	check(err)
	return results
}

// Given a bunch of `RhymebrainResult`s, find words with the highest score.
// For example, if there are 2 words with a score of 200 and 3 with a score of
// 300, it will return the 3 with a score of 300.
func wordsWithMaximumScore(rhymeBrainResults []RhymebrainResult) []string {
	words := make([]string, 0)
	maxScore := rhymeBrainResults[0].Score

	for _, rhymeBrainResult := range rhymeBrainResults {
		if rhymeBrainResult.Score > maxScore {
			maxScore = rhymeBrainResult.Score
		}
	}

	for _, rhymeBrainResult := range rhymeBrainResults {
		if rhymeBrainResult.Score == maxScore {
			words = append(words, rhymeBrainResult.Word)
		}
	}

	return words
}

func main() {
	// 1) Iterate through each phrase and rhyme
	// 2) Find phrases that have a rhyme in them
	// 3) Replace the rhyme with the original word in the phrase
	// 4) Collect the changed phrases in a slice
	// 5) Print out the changed phrases

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: go run Gorls.go [word to pun on]\n")
		os.Exit(64)
	}
	keyword := os.Args[1]
	fmt.Printf("Punning on \"%s\"\n", keyword)
	puns := []Pun{}
	rhymes := wordsWithMaximumScore(rhymes(keyword))

	phraseFilePaths := []string{
		"phrases/beatles_songs.txt",
		"phrases/best-selling-books.txt",
		"phrases/movie-quotes.txt",
		"phrases/oscar_winning_movies.txt",
		"phrases/wikipedia_idioms.txt",
	}

	for _, path := range phraseFilePaths {
		contents, err := ioutil.ReadFile(path)
		check(err)
		phrases := strings.Split(string(contents), "\n")

		for _, rhyme := range rhymes {
			re := regexp.MustCompile("(?i)\\b" + rhyme + "\\b")
			for _, phrase := range phrases {
				if re.MatchString(phrase) {
					// Replace the rhyme with the word we're punning on.
					// If the word is "carts", Lonely Hearts -> Lonely Carts
					punPhrase := re.ReplaceAllString(phrase, keyword)
					pun := Pun{PunPhrase: punPhrase, OriginalPhrase: phrase}
					puns = append(puns, pun)
				}
			}
		}
	}

	if len(puns) == 0 {
		fmt.Println("No puns found :(")
	} else {
		for _, pun := range puns {
			fmt.Println(pun.PunPhrase + " (pun of " + pun.OriginalPhrase + ")")
		}
	}
}
