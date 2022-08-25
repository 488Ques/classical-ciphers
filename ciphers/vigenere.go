package ciphers

import "strings"

type Vigenere struct {
	alphabet string
	key      string
	msg      string
}

func (vig *Vigenere) Parse(data []byte) error {
	splitted := strings.Split(string(data), "\n")

	if len(splitted) != 3 {
		return errInvalidInputFormat
	}

	vig.alphabet = splitted[0]
	vig.key = splitted[1]
	vig.msg = splitted[2]

	return nil
}

func (vig *Vigenere) Encrypt() (string, error) {
	alphabetIndices := indexAlphabet(vig.alphabet)

	cipherText := make([]byte, len(vig.msg))
	keyLen := len(vig.key)
	alphabetLen := len(vig.alphabet)
	j := 0
	for i, r := range vig.msg {
		if r == ' ' {
			cipherText[i] = ' '
			continue
		}
		msgIndex, ok := alphabetIndices[r]
		if !ok {
			return "", errIllegalCharInMsg(r)
		}
		keyIndex, ok := alphabetIndices[rune(vig.key[j%keyLen])]
		if !ok {
			return "", errIllegalCharInMsg(r)
		}
		encryptedIndex := (msgIndex + keyIndex) % alphabetLen
		cipherText[i] = vig.alphabet[encryptedIndex]
		j++
	}

	return string(cipherText), nil
}
