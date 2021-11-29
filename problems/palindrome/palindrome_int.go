package palindrome

import "fmt"

func IsNumberPalindrome(num int32) {
	var mod int32
	var tempNum = num
	for {
		tmp := tempNum % 10
		if tmp != 0 {
			mod = (mod * 10) + tmp
			tempNum = tempNum / 10
		} else {
			break
		}
	}

	if num == mod {
		fmt.Println("PALINDROME")
	} else {
		fmt.Println("NOT PALINDROME")
	}
}
