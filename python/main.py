"""Determine the sum of all multiples of 3 or 5 up to 1000."""

def is_multiple(number):
    """return True if the number is a multiple of 3 or 5 else return False"""
    return not (number % 3 and number % 5)

def sum_multiples(upper_limit):
    """
    Sum across all multiples of 3 or 5 up to but not including the upper
    limit
    """
    return sum([i for i in range(upper_limit) if is_multiple(i)])

if __name__ == "__main__":
    print sum_multiples(1000)
