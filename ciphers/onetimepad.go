package ciphers

import "strings"

type OTP struct {
	alphabet string
	key      string
	msg      string
}

func (otp *OTP) Parse(data []byte) error {
	splitted := strings.Split(string(data), "\n")

	if len(splitted) != 3 {
		return errInvalidInputFormat
	}

	if len(splitted[0]) <= 0 || len(splitted[1]) <= 0 || len(splitted[2]) <= 0 {
		return errInvalidInputFormat
	}

	if len(splitted[1]) != len(splitted[2]) {
		return errPlainAndCipherNotEqual
	}

	otp.alphabet = splitted[0]
	otp.key = splitted[1]
	otp.msg = splitted[2]

	return nil
}

func (otp *OTP) Encrypt() (string, error) {
	alphabetIndices := indexAlphabet(otp.alphabet)

	cipherText := make([]rune, len(otp.msg))
	alphabetLen := len(otp.alphabet)
	for i, r := range otp.msg {
		if r == ' ' {
			cipherText[i] = ' '
			continue
		}

		msgIndex, ok := alphabetIndices[r]
		if !ok {
			return "", errIllegalCharInMsg(r)
		}
		keyIndex, ok := alphabetIndices[rune(otp.key[i])]
		if !ok {
			return "", errIllegalCharInMsg(r)
		}
		cipherIndex := (msgIndex + keyIndex) % alphabetLen
		cipherText[i] = rune(otp.alphabet[cipherIndex])
	}

	return string(cipherText), nil
}
