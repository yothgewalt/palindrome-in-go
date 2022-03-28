package function

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type signedOrString interface {
	constraints.Signed | string
}

func palindrome[input signedOrString](sliceInput []input) bool {
	length := len(sliceInput)
	for character := 0; character < length/2; character++ {
		if sliceInput[character] != sliceInput[length-1-character] {
			return false
		}
	}

	return true
}

func sliceInteger[signed constraints.Signed](input signed) (slice []int) {
	for input > 0 {
		slice = append(slice, int(input%10))
		input /= 10
	}

	return
}

func IsPalindrome[input signedOrString](firstInput input) bool {
	return func(toInterface interface{}) bool {
		var sliceInt []int
		switch newType := toInterface.(type) {
		case string:
			return palindrome(strings.Split(newType, ""))
		case int:
			sliceInt = sliceInteger(newType)
		case int8:
			sliceInt = sliceInteger(newType)
		case int16:
			sliceInt = sliceInteger(newType)
		case int32:
			sliceInt = sliceInteger(newType)
		case int64:
			sliceInt = sliceInteger(newType)
		}

		return palindrome(sliceInt)
	}(firstInput)
}
