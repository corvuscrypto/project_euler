#include <stdio.h>

/**
 * determine if a number is a multiple of 3 or 5
 * @param  number integer to check
 * @return        1 if true, 0 otherwise
 */
int isMultiple(int number) {
  return !(number % 3 && number % 5);
}

/**
 * sum accross all multiples of 3 or 5 under the upper integer limit given
 * @param  numberArray the integer at which we stop summing (non-inclusive)
 * @return             sum of the numbers which are multiples of 3 or 5
 */
int sumMultiples(int upperLimit) {
  int sum = 0;
  for (int i = 0; i < upperLimit; ++i) {
    if (isMultiple(i)) sum += i;
  }
  return sum;
}

int main() {
  //generate the sum according to the Euler problem in the readme
  printf("The sum of all integers below 1000 is: %d\n", sumMultiples(1000));
  return 0;
}
