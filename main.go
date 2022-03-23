package main

import (
	"fmt"
	"regexp"
	"strconv"
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

func main() {
	fmt.Println(testValidity("23-ab-48-caba-56-abc"))
	fmt.Println(averageNumber("23-ab-48-caba-56-abc"))
}
