package main

import (
	"fmt"
	"strings"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go stringSender(c1)
	go stringSpliter(c1, c2)
	printGopher(c2)
}

func printGopher(ch2 chan string) {
	for _, v := range <-ch2 {
		fmt.Println(v)
	}
}

func stringSender(ch1 chan string) {
	for _, v := range []string{"what the fuck", "fuck that and fuck this"} {
		ch1 <- v
	}
	close((ch1))
}

func stringSpliter(ch1 chan string, ch2 chan string) {
	for {
		str, ok := <-ch1
		if !ok {
			close(ch2)
			return
		}
		words := strings.Fields(str)
		fmt.Println(words)

		for _, w := range words {
			ch2 <- w
		}
		close(ch2)
	}

}
