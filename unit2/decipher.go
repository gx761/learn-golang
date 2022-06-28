package main

import "fmt"

func main() {

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
