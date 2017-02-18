package main

import (
	"fmt"
	"math"
)

//isPalindrome determines whether a number is a palindrome or not
func isPalindrome(num int) bool {
	// use a mod-10 shift method

	reversedNum := 0
	for tempNum := num; tempNum > 0; tempNum /= 10 {
		reversedNum *= 10
		reversedNum += tempNum % 10
	}
	//now do equality check
	return reversedNum == num
}

//getLargestPalindrome crawls through products of 3-digit numbers from the top down
//until we find the largest palindrome (by comparison update)
func getLargestPalindrome(numDigits int) (product int) {
	startLimit := int(math.Pow(10, float64(numDigits))) - 1
	for j := startLimit; j > 0; j-- {
		for k := startLimit; k >= j; k-- {
			if newProduct := j * k; isPalindrome(newProduct) {
				if newProduct > product {
					product = newProduct
				}
			}
		}
	}
	return product
}

func main() {
	fmt.Println(getLargestPalindrome(3))
}
