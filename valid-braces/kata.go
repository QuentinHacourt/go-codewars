package kata

var brackets = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func ValidBraces(str string) bool {
	var openBraces []string

	for _, r := range str {
		letter := string(r)

		if len(openBraces) > 0 && brackets[openBraces[len(openBraces)-1]] == letter {
			openBraces = openBraces[:len(openBraces)-1]
		} else {
			openBraces = append(openBraces, letter)
		}
	}

	return len(openBraces) == 0
}
