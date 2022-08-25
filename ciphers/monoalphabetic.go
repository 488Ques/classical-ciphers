package ciphers

import (
	"strings"
)

type Monoalphabetic struct {
	alphabet string
	key      string
	msg      string
}

func (mono *Monoalphabetic) Parse(data []byte) error {
	splitted := strings.Split(string(data), "\n")

	if len(splitted) != 3 {
		return errInvalidInputFormat
	}

	if len(splitted[0]) <= 0 || len(splitted[1]) <= 0 || len(splitted[2]) <= 0 {
		return errInvalidInputFormat
	}

	if len(splitted[0]) != len(splitted[1]) {
		return errPlainAndCipherNotEqual
	}

	mono.alphabet = splitted[0]
	mono.key = splitted[1]
	mono.msg = splitted[2]

	return nil
}

func (mono *Monoalphabetic) Encrypt() (string, error) {
	cipherText := make([]rune, len(mono.msg))
	lettersMap := mapLetters(mono.alphabet, mono.key)
	if _, ok := lettersMap[' ']; !ok {
		lettersMap[' '] = ' '
	}

	for i, r := range mono.msg {
		if _, ok := lettersMap[r]; !ok {
			return "", errIllegalCharInMsg(r)
		}
		cipherText[i] = lettersMap[r]
	}

	return string(cipherText), nil
}
