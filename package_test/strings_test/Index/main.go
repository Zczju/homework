package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"
	substring := "ll"
	firstOccur := strings.Index(s, substring)
	fmt.Println(firstOccur)
}
