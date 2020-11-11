package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println(quote.Go())

	handleRequests()
}
