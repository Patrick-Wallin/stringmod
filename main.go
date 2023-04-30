package main

import (
	"fmt"

	"github.com/patrick-wallin/stringmod/quotes"
	"github.com/patrick-wallin/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(quotes.GetEmoji("turtle"))
}
