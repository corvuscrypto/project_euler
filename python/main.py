"""Calculate biggest prime factor"""


def is_prime(num):
    """Detects whether a number is prime based on a slow and simple heuristic"""
    if num <= 1:
        return False
    if num <= 3:
        return True
    if num % 2 == 0 or num % 3 == 0:
        return False
    _num = 5
    while (_num ** 2) <= num:
        if num % _num == 0 or num % (_num + 2) == 0:
            return False
        _num += 6
    return True


def get_biggest_prime_factor(num):
    """Loops from sqrt of num down to 0 to find biggest prime factor."""
    _num = int(num ** 0.5)
    while _num > 0:
        if num % _num == 0:
            if is_prime(_num):
                return _num
        _num -= 1
    return -1


if __name__ == "__main__":
    print get_biggest_prime_factor(600851475143)
