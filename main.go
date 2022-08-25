package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/488Ques/classical-ciphers/ciphers"
	"golang.org/x/exp/maps"
)

type Cipher interface {
	Parse([]byte) error
	Encrypt() (string, error)
}

var ciphersMap = map[string]Cipher{
	"monoalphabetic": &ciphers.Monoalphabetic{},
	"vigenere":       &ciphers.Vigenere{},
	"otp":            &ciphers.OTP{},
	"railfence":      &ciphers.Railfence{},
}

func main() {
	fileDir := flag.String("file", "./input", "Input file directory")
	cipherName := flag.String("cipher", "", func() string {
		return fmt.Sprint("Name of the cipher. Must match one of these: ", maps.Keys(ciphersMap))
	}())
	flag.Parse()

	if len(*fileDir) == 0 {
		log.Fatal("Input file directory must not be empty string")
	}
	if len(*cipherName) == 0 {
		log.Fatal("Cipher must not be empty string")
	}

	cipherText, err := Execute(*fileDir, *cipherName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cipherText)
}

func Execute(fileDir, cipherName string) (string, error) {
	data, err := os.ReadFile(fileDir)
	if err != nil {
		return "", err
	}

	cipher, ok := ciphersMap[cipherName]
	if !ok {
		return "", errors.New("Invalid cipher name")
	}

	err = cipher.Parse(data)
	if err != nil {
		return "", err
	}

	cipherText, err := cipher.Encrypt()
	if err != nil {
		return "", err
	}

	return cipherText, nil
}
