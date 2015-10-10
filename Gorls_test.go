package main

import (
	"testing"
)

func testEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestWordsWithMaximumScore(t *testing.T) {
	s := []RhymebrainResult{}
	s = append(s, RhymebrainResult{"one", 300})
	s = append(s, RhymebrainResult{"two", 300})
	s = append(s, RhymebrainResult{"three", 200})

	r := wordsWithMaximumScore(s)
	e := []string{"one", "two"}

	if !testEq(r, e) {
		t.Errorf("wordsWithMaximumScore() = %v want %v", r, e)
	}
}
