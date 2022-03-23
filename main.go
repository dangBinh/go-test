package main

import (
	"fmt"
	"math"
	"math/rand"
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
	shortestWord := FindShortestWord(aWords)
	longestWord := FindLongestWord(aWords)
	avgWordLength := CalcAverageWordLength(aWords)
	aStr := ListOfWordLengthEqualAverage(aWords, avgWordLength)

	return shortestWord, longestWord, avgWordLength, aStr
}

func StringSeeder(length int) string {
	b := make([]byte, length)
	const charset = "abcdefghijklmnopqrstuvwxyz"

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// The difficulty of the function: O(n)
func generateString(flag bool) string {
	var testStr string
	// Random correct string
	if flag {
		var aStr []string

		numberOfPattern := rand.Intn(50)
		for i := 1; i <= numberOfPattern; i++ {
			str := strconv.Itoa(rand.Intn(100)) + "-" + StringSeeder(10)
			aStr = append(aStr, str)
		}
		testStr = strings.Join(aStr, "-")
	} else {
		testStr = StringSeeder(100)
	}

	return testStr
}

func main() {
	text := "23-ab-48-caba-56-abc"
	fmt.Println(testValidity(text))
	fmt.Println(averageNumber(text))
	fmt.Println(wholeStory(text))
	fmt.Println(storyStats(text))
	generateString(true)
}
