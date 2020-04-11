package main

import (
	"fmt"

	"github.com/irajdeep/go-sandbox/gobook/ch7/7.1/wordcount"
)

func main() {

	var b wordcount.ByteWordsCounter
	str := "a  quick   brown fox just jumped   without  a lazy dog"

	fmt.Fprintf(&b, str)
	fmt.Println(b)
}
