package main

import "fmt"

//IsMultiple checks whether the given number is a multiple of 3 or 5
func IsMultiple(number int) bool {
	return number%3 == 0 || number%5 == 0
}

//SumMultiples sums across all multiples of 3 and 5 up to the limit specified
func SumMultiples(upperLimit int) (sum int) {
	for i := 0; i < upperLimit; i++ {
		if IsMultiple(i) {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(SumMultiples(1000))
}
