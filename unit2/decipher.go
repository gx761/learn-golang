package main

import (
	"fmt"
	"strings"
)

func main() {
	decipher()
	cipher()

}

func cipher() {
	plainText := "your message goes here"
	keyword := "GOLANG"
	tripAndUppered := strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	var ciphered string
	for i, c := range tripAndUppered {
		key := rune(keyword[i%len(keyword)])
		shift := key - 'A'

		if c+shift <= 'Z' {
			ciphered = ciphered + string(c+shift)
		} else {
			ciphered = ciphered + string(c-26+shift)
		}
	}
	fmt.Println(ciphered)

}

func decipher() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	var deciphered string

	for i, c := range cipherText {
		key := rune(keyword[i%len(keyword)])
		shift := key - 'A'

		if c-shift >= 'A' {
			deciphered = deciphered + string(c-shift)
		} else {
			deciphered = deciphered + string(c+26-shift)
		}
	}
	fmt.Println(deciphered)
}
