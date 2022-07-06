package main

import (
	"fmt"
	"net/url"
)

func parseUrl(u string) {
	i, err := url.Parse(u)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T, %#v\n", err, err)
		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("Url:", e.URL)
			fmt.Println("Err:", e.Err)
		}
	} else {
		fmt.Printf("%#v", i)
	}
}

func main() {
	parseUrl("https://www.bai du.com/test")
}
