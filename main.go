package main

import (
	"fmt"

	"github.com/Patrick-Wallin/stringmod/quotes"
	"github.com/Patrick-Wallin/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(quotes.GetEmoji("turtle"))
}
