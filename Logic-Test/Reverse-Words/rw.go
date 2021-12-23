package main

import (
	"fmt"
)

func main() {
	rev := ReverseWord("I am A Great human")
	fmt.Println(rev)
}

func ReverseWord(kata string) string {
	var reverse string
	for _, v := range kata {
		reverse = string(v) + reverse
	}
	return reverse
}
