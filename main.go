package main

import (
	"fmt"

	"github.com/yongyuth-chuankhuntod/palindrome-in-go/function"
)

func main() {
	fmt.Println(function.IsPalindrome("reviver"))    /* true */
	fmt.Println(function.IsPalindrome(20211202))     /* true */
	fmt.Println(function.IsPalindrome("palindrome")) /* false */
	fmt.Println(function.IsPalindrome(3282022))      /* false */
}
