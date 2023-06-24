package main

import (
	"fmt"
	"github.com/fgsoftware1/reverse-sentence"
)

func main() {
	sentence := "Hello world!"
	reversed := reverselib.Reverse(sentence)
	fmt.Println(reversed)
}