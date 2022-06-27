package main

import (
	"fmt"
)


func main() {
	ciphered:= "L fdph, L vdz, L frqtxhuhg."
	var deciphered string

	for _, c:= range ciphered {

		if((c>='d'&&c<'z')) {
			c = c-3
		} else if(c>='a' && c<'d') {
			c = c + 23
		}

		if((c>='D'&&c<'Z')) {
			c = c-3
		} else if(c>='A' && c<'D') {
			c = c + 23
		}
		deciphered = deciphered + string(c)
	}
	fmt.Print(deciphered)
}