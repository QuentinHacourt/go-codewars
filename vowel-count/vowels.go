package kata

func GetCount(str string) (count int) {
	count = 0

	for _, character := range str {
		letter := string(character)
		if isVowel(letter) {
			count++
		}

	}

	return count
}

func isVowel(letter string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, vowel := range vowels {
		if vowel == letter {
			return true
		}
	}

	return false
}
