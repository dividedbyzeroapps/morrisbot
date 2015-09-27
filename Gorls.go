package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func rhymes(word string) []RhymebrainResult {
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
	// TODO:
	// 1) Iterate through all phrases and all rhymes,
	// 2) find phrases that have a rhyme in them
	// 3) replace the rhyme with the keyword in the phrase
	// 4) collect the changed phrases in a slice
	// 5) print out the changed phrases

	keyword := "coin"
	puns := []string{}
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
				matches := re.MatchString(phrase)
				if matches {
					// Replace the rhyme with the word we're punning on.
					// If the word is "carts", Lonely Hearts -> Lonely Carts
					pun := re.ReplaceAllString(phrase, keyword)
					puns = append(puns, pun)
				}
			}
		}
	}

	for _, pun := range puns {
		fmt.Println(pun)
	}
}
