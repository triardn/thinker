package palindrome

import "fmt"

func IsStringPalindrome(str string) {
	var strReverse string

	for _, word := range str {
		strReverse = string(word) + strReverse
	}

	if str == strReverse {
		fmt.Println("PALINDROME")
	} else {
		fmt.Println("NOT PALINDROME")
	}
}
