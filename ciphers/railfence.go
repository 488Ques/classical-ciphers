package ciphers

import (
	"strconv"
	"strings"
)

type Railfence struct {
	key int
	msg string
}

func (rf *Railfence) Parse(data []byte) error {
	splitted := strings.Split(string(data), "\n")

	if len(splitted) != 2 {
		return errInvalidInputFormat
	}

	if len(splitted[0]) <= 0 || len(splitted[1]) <= 0 {
		return errInvalidInputFormat
	}

	key, err := strconv.ParseInt(splitted[0], 10, 32)
	if err != nil {
		return err
	}
	rf.key = int(key)
	rf.msg = splitted[1]

	return nil
}

// TODO ignore whitespace
func (rf *Railfence) Encrypt() (string, error) {
	cols := rf.key
	rows := len(rf.msg) / cols
	if cells := rows * cols; cells != len(rf.msg) {
		rows++
	}

	table := make([][]rune, rows)
	for i := range table {
		table[i] = make([]rune, cols)
	}

	msgIndex := 0
	for col := range table {
		for row := range table[col] {
			table[col][row] = rune(rf.msg[msgIndex])
			msgIndex++
		}
	}

	msgIndex = 0
	cipherText := make([]rune, len(rf.msg))
	for col := range table[0] {
		for row := range table {
			cipherText[msgIndex] = table[row][col]
			msgIndex++
		}
	}

	return string(cipherText), nil
}
