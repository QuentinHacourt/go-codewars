package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Word: din\nResult: %v\n", DuplicateEncode("din"))
	fmt.Printf("Word: recede\nResult: %v\n", DuplicateEncode("recede"))
	fmt.Printf("Word: Success\nResult: %v\n", DuplicateEncode("Success"))
	fmt.Printf("Word: (( @\nResult: %v\n", DuplicateEncode("(( @"))

}

func DuplicateEncode(word string) string {
	chars := make(map[string]string)

	for _, char := range word {
		lowerChar := strings.ToLower(string(char))
		if _, ok := chars[lowerChar]; ok {
			chars[string(lowerChar)] = ")"
		} else {
			chars[string(lowerChar)] = "("
		}
	}

	res := ""

	for _, char := range word {
		lowerChar := strings.ToLower(string(char))
		res += chars[string(lowerChar)]
	}

	return res
}
