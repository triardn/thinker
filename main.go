package main

import (
	"fmt"

	missingnumber "github.com/triardn/websocket-example/thinker/problems/missing_number"
	"github.com/triardn/websocket-example/thinker/problems/palindrome"
)

func main() {
	palindrome.IsStringPalindrome("katak")
	palindrome.IsNumberPalindrome(1221)
	missingNumber := missingnumber.FindMissingNumber([]int{0, 3, 1})
	fmt.Println(missingNumber)
}
