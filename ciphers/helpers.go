package ciphers

func mapLetters(alphabet, key string) map[rune]rune {
	lettersMap := make(map[rune]rune)
	for i, r := range alphabet {
		lettersMap[r] = rune(key[i])
	}

	return lettersMap
}

func indexAlphabet(alphabet string) map[rune]int {
	alphabetIndices := make(map[rune]int)
	for i, c := range alphabet {
		alphabetIndices[c] = i
	}
	return alphabetIndices
}
