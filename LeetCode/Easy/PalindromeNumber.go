package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := fmt.Sprint(x)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}
