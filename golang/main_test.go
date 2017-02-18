package main

import "testing"

func TestIsPrime(T *testing.T) {

	trueValues := []int{3, 5, 7}
	falseValues := []int{1, 8, 10, 10000}

	for _, input := range trueValues {
		if isPrime(input) == false {
			T.Error(input)
		}
	}

	for _, input := range falseValues {
		if isPrime(input) == true {
			T.Error(input)
		}
	}

}

func TestGetBiggestPrimeFactor(T *testing.T) {
	if getBiggestPrimeFactor(13195) != 29 {
		T.Fail()
	}
}
