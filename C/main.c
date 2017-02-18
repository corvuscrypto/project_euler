#include <stdio.h>
#include <math.h>

int is_palindrome(int num) {
  //use mod-10 shift method
  int temp_num;
  int reversed_num = 0;
  for (temp_num = num; temp_num > 0; temp_num /= 10) {
    reversed_num *= 10;
    reversed_num += temp_num % 10;
  }
  return reversed_num == num;
}

int get_largest_palindrome(int num_digits) {
  int start_limit = (int)(pow(10, num_digits)) - 1;
  int product = 0;
  for(int j = start_limit; j > 0; --j) {
    for (int k = start_limit; k >= j; --k) {
      int new_product = k * j;
      if (is_palindrome(new_product) && new_product > product) {
        product = new_product;
      }
    }
  }
  return product;
}

int main(){
  //mini test (again built into the program because why not)
  if (get_largest_palindrome(2) != 9009) {
    puts("Method is failing!");
    return 0;
  }
  printf("%d\n", get_largest_palindrome(3));
}
