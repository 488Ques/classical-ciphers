package ciphers

import (
	"errors"
	"fmt"
)

var (
	errInvalidInputFormat = errors.New("Input's format is invalid")

	errPlainAndCipherNotEqual = errors.New("Plaintext's length and ciphertext's are not equal")
)

func errIllegalCharInMsg(char rune) error {
	return fmt.Errorf("Message contains illegal character (%q)", char)
}
