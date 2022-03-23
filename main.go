package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// The difficulty of the function: O(n)
// testValidity checks whether text complies with the format.
func testValidity(text string) bool {
	regex := `^(([0-9]+[-][a-z]+)(?:[-]*)?)+$`
	if m, _ := regexp.MatchString(regex, text); !m {
		return false
	}
	return true
}

// The difficulty of the function: O(n)
// averageNumber calculates the average number from all the numbers
func averageNumber(text string) (avgNumber float32) {
	re := regexp.MustCompile(`[0-9]+`)
	matches := re.FindAllString(text, -1)
	length := len(matches)

	if length == 0 {
		return 0
	}
	var total float32
	for _, match := range matches {
		fmt.Println(match)
		number, _ := strconv.ParseFloat(match, 10)
		total += float32(number)
	}
	return total / float32(length)
}

// The difficulty of the function: O(n)
// wholeStory returns text that composed from all words
func wholeStory(text string) (story string) {
	re := regexp.MustCompile(`[a-z]+`)
	matches := re.FindAllString(text, -1)

	return strings.Join(matches, " ")
}

func FindShortestWord(words []string) int {
	if len(words) == 0 {
		return 0
	}

	shortest := len(words[0])
	for _, word := range words {
		if len(word) < shortest {
			shortest = len(word)
		}
	}

	return shortest
}

func FindLongestWord(words []string) int {
	if len(words) == 0 {
		return 0
	}

	longest := len(words[0])
	for _, word := range words {
		if len(word) > longest {
			longest = len(word)
		}
	}

	return longest
}

func CalcAverageWordLength(words []string) float32 {

	if len(words) == 0 {
		return 0
	}

	var total float32
	for _, word := range words {
		total += float32(len(word))
	}
	return total / float32(len(words))
}

func ListOfWordLengthEqualAverage(aWords []string, avgWordLength float32) []string {
	var words []string

	for _, word := range aWords {
		if len(word) == int(math.Round(float64(avgWordLength))) {
			words = append(words, word)
		}
	}

	return words
}

// The difficulty of the function: O(n)
//  storyStats returns stats of text
func storyStats(text string) (int, int, float32, []string) {
	words := wholeStory(text)
	aWords := strings.Split(words, " ")
	fmt.Println(aWords)
	shortestWord := FindShortestWord(aWords)
	longestWord := FindLongestWord(aWords)
	avgWordLength := CalcAverageWordLength(aWords)
	aStr := ListOfWordLengthEqualAverage(aWords, avgWordLength)

	return shortestWord, longestWord, avgWordLength, aStr
}

func main() {
	fmt.Println(testValidity("23-ab-48-caba-56-abc"))
	fmt.Println(averageNumber("23-ab-48-caba-56-abc"))
	fmt.Println(wholeStory("23-ab-48-caba-56-abc"))
	fmt.Println(storyStats("23-ab-48-caba-56-abc"))
}
