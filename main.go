package main

import (
	"fmt"
	"regexp"
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

func main() {
	fmt.Println(testValidity("23-ab-48-caba-56-abc"))
}
