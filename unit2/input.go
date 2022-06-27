package main

import (
	"errors"
	"fmt"
)

func main() {
	value := "sdf"

	r := convert(value)

	if r.err != nil {
		fmt.Printf("%v, %T", r.err, r.err)
	} else {
		fmt.Printf("%v, %T", r.v, r.v)
	}
}

type result struct {
	v   bool
	err error
}

func convert(v string) result {
	switch v {
	case "true", "yes", "1":
		return result{true, nil}
	case "false", "no", "0":
		return result{false, nil}
	default:
		return result{true, errors.New("input is invalid")}
	}
}
