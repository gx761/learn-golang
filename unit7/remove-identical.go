package main

import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go msgSender(c1)
	go msgFilter(c1, c2)
	printGopher(c2)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func msgSender(ch chan string) {
	// fmt.Println("start running msg sender")
	for _, v := range []string{"123", "456", "789", "456", "456", "879"} {
		ch <- v
	}
	close(ch)
}

func msgFilter(ch1 chan string, ch2 chan string) {
	// fmt.Println("start running msg filter")
	prev := ""
	for {
		msg, ok := <-ch1
		if !ok {
			// fmt.Println("channel closed")
			close(ch2)
			return
		}
		// fmt.Println("msg is", msg)
		if msg != prev {
			ch2 <- msg
			prev = msg
			// fmt.Println("msg is", msg)
		} else {
			// fmt.Println("msg is same as previous msg")
		}
	}
}
