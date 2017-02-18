"""Test for the main program functions"""
import main


def test_is_prime():
    """Test for the is_prime function"""
    assert main.is_prime(3) is True
    assert main.is_prime(5) is True
    assert main.is_prime(7) is True
    assert main.is_prime(8) is False
    assert main.is_prime(100) is False
    assert main.is_prime(100000000) is False


def test_get_biggest_prime_factor():
    """Test for the get_biggest_prime_factor function"""
    assert main.get_biggest_prime_factor(13195) == 29
