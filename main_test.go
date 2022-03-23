package main

import (
	"testing"
)

func TestTestValidity(t *testing.T) {
	validString := generateString(true)

	if !testValidity(validString) {
		t.Error("Failed!")
	}

	inValidString := generateString(true)

	if !testValidity(inValidString) {
		t.Error("Failed!")
	}
}

func TestAverageNumber(t *testing.T) {
	validString := generateString(true)

	if len(validString) > 0 && averageNumber(validString) == 0 {
		t.Error("Failed!")
	}
}

func TestWholeStory(t *testing.T) {
	validString := generateString(true)

	if len(validString) > 0 && len(wholeStory(validString)) == 0 {
		t.Error("Failed!")
	}
}

func TestStoryStats(t *testing.T) {

	validString := generateString(true)
	shortestWord, _, _, _ := storyStats(validString)
	if len(validString) > 0 && shortestWord == 0 {
		t.Error("Failed!")
	}

}
