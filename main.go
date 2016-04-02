package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type RhymebrainResult struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

type Pun struct {
	Keyword        string `json:,omitempty`
	PunPhrase      string `json:,omitempty`
	OriginalPhrase string `json:,omitempty`
	Error          string `json:,omitempty`
}

func rhymes(word string) []RhymebrainResult {
	url := "http://rhymebrain.com/talk?function=getRhymes&word=" + word + "&maxResults=0&lang=en"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	rawJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	results := make([]RhymebrainResult, 0)
	err = json.Unmarshal(rawJSON, &results)
	if err != nil {
		panic(err)
	}

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

func sample(orig []Pun) Pun {
	return orig[rand.Int()%len(orig)]
}

func loadPhrases() []string {
	phraseFilePaths := []string{
		"phrases/beatles_songs.txt",
		"phrases/best-selling-books.txt",
		"phrases/movie-quotes.txt",
		"phrases/oscar_winning_movies.txt",
		"phrases/wikipedia_idioms.txt",
	}

	var phrases []string

	for _, path := range phraseFilePaths {
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		phrases = append(phrases, strings.Split(string(contents), "\n")...)
	}

	return phrases
}

func randomPun(phrases []string, keyword string) Pun {
	puns := []Pun{}
	rhymes := wordsWithMaximumScore(rhymes(keyword))
	for _, rhyme := range rhymes {
		re := regexp.MustCompile("(?i)\\b" + rhyme + "\\b")
		for _, phrase := range phrases {
			if re.MatchString(phrase) {
				// Replace the rhyme with the word we're punning on.
				// If the word is "carts", Lonely Hearts -> Lonely Carts
				punPhrase := re.ReplaceAllString(phrase, keyword)
				puns = append(puns, Pun{Keyword: keyword, PunPhrase: punPhrase, OriginalPhrase: phrase})
			}
		}
	}

	if len(puns) == 0 {
		return Pun{Keyword: keyword, Error: "No puns found :("}
	} else {
		return sample(puns)
	}
}

func main() {
	phrases := loadPhrases()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pun := randomPun(phrases, r.URL.Path[1:])
		resp, err := json.MarshalIndent(pun, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, string(resp))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
