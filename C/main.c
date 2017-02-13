#include <stdio.h>

/**
 * perform summation on even fibonnaci numbers up to the limit using a pointer
 * to a storage variable for optimized space utilization
 * @param alpha the term before the last
 * @param beta  the last term
 * @param limit the numerical limit
 * @param sum   pointer to the sum variable
 */
void generateFibonacciSum(int alpha, int beta, int limit, int *sum) {
  //implement with tail call and no new declarations for optimized recursion on
  // -O2 flag
  alpha += beta;
  if (beta % 2 == 0) {
    *sum += beta;
  }
  if (alpha > limit) {
    return;
  }
  generateFibonacciSum(beta, alpha, limit, sum);
}

int main() {
  //generate the sum according to the Euler problem in the readme
  // make a holding variable
  int sum = 0;
  //perform the summation starting at [1,2]
  generateFibonacciSum(1, 2, 4e6, &sum);
  printf("The sum of all even fibonnacis under 4,000,000: %d\n", sum);
  return 0;
}
