package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	//do inline checks
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for _num := 5; _num*_num <= num; _num += 6 {
		if num%_num == 0 || num%(_num+2) == 0 {
			return false
		}
	}
	return true
}

func getBiggestPrimeFactor(num int) int {

	for _num := int(math.Sqrt(float64(num))); _num*_num <= num; _num-- {
		if num%_num == 0 {
			if isPrime(_num) {
				return _num
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(getBiggestPrimeFactor(600851475143))
}
