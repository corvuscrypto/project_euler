package main

import "fmt"

//FibonacciSum performs a sum on all even fibonnaci numbers
func FibonacciSum(alpha, beta, limit int, sum *int) {
	alpha += beta
	if beta%2 == 0 {
		*sum += beta
	}
	if alpha > limit {
		return
	}
	FibonacciSum(beta, alpha, limit, sum)
}

func main() {
	//make a holding var
	sum := 0
	FibonacciSum(1, 2, 4e6, &sum)
	fmt.Println("The sum of all even fibonnacis up to 4e6 is: ", sum)
}
