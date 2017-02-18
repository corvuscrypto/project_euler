package main

import "testing"

func TestIsPalindrome(T *testing.T) {
	palindromes := []int{33533, 222222, 66455466}
	nonPalindromes := []int{665, 244, 188888888, 10}
	for _, input := range palindromes {
		if isPalindrome(input) == false {
			T.Error(input)
		}
	}
	for _, input := range nonPalindromes {
		if isPalindrome(input) == true {
			T.Error(input)
		}
	}
}

func TestGetLargestPalindrome(T *testing.T) {
	//since we are given the palindrome for 2 digits, our generalized approach should work and give
	//the expected answer
	if value := getLargestPalindrome(2); value != 9009 {
		T.Error(value, "Approach failed")

	}
}
