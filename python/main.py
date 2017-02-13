"""Determine sum of all even fibonnaci numbers up to 4,000,000"""

# declare global sum variable
sum_holder = 0

def fibonnaci_sum(alpha, beta, limit):
    """perform tail-call recursive fibbonaci sequence and adjust the sum"""
    global sum_holder
    alpha += beta
    # if beta is even then add it to sum_holder
    if beta % 2 == 0:
        sum_holder += beta
    if alpha > limit:
        return
    fibonnaci_sum(beta, alpha, limit)

if __name__ == "__main__":
    fibonnaci_sum(1, 2, 4e6)
    print sum_holder
