#include <stdio.h>
#include <math.h>

/**
 * Determines if a number given is prime
 * @param  num the number being tested
 * @return     bool (1 if prime else 0)
 */
int is_prime(long int num) {
  //do some inline checks before looping
  if (num <= 1) return 0;
  if (num <= 3) return 1;
  if (num % 2 == 0 || num % 3 == 0) return 0;
  long int _num;
  // loop only from the square root of the number (product factors above not necessary to compute)
  for (_num = 5; _num * _num <= num; _num += 6) {
    if (num % _num == 0) return 0;
    if (num % (_num + 2) == 0) return 0;
  }
  return 1;
}

long int biggest_prime_factor(long int num) {
  long int i;
  for (i = (long int)sqrt(num); i > 0; --i) {
    if (num % i == 0) {
      if (is_prime(i)) return i;
    }
  }
  return -1;
}

int main() {
  // perform a mini test (we expect that f(13195) == 29)
  if (biggest_prime_factor(13195) != 29) {
    puts("Function incorrectly working, git gud");
    printf("%ld\n", biggest_prime_factor(13195));
    return 0;
  }
  printf("Answer is: %ld\n", biggest_prime_factor(600851475143));
  return 0;
}
