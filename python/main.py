"""find the largest palindrome"""


def is_palindrome(num):
    reversed_num = 0
    temp_num = num
    while temp_num > 0:
        reversed_num *= 10
        reversed_num += temp_num % 10
        temp_num /= 10
    return reversed_num == num


def get_largest_palindrome(num_digits):
    start_limit = int(10 ** num_digits)
    product = 0
    for j in range(start_limit)[::-1]:
        for k in range(j, start_limit)[::-1]:
            new_product = j * k
            if is_palindrome(new_product) and new_product > product:
                product = new_product
    return product


if __name__ == "__main__":
    print get_largest_palindrome(3)
