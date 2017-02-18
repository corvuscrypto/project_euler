"""find the largest palindrome"""
import main


def test_is_palindrome():
    """tests the palindrome check function"""
    for _input in [33, 565, 77733777]:
        assert main.is_palindrome(_input) is True
    for _input in [334, 56, 7773377]:
        assert main.is_palindrome(_input) is False


def test_get_largest_palindrome():
    """tests the function that gets the largest palindrome"""
    assert main.get_largest_palindrome(2) == 9009
