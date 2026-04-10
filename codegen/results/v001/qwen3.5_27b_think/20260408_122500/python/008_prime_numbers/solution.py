import sys
import math

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    limit = int(math.sqrt(n)) + 1
    for i in range(3, limit, 2):
        if n % i == 0:
            return False
    return True

if len(sys.argv) == 2:
    try:
        N = int(sys.argv[1])
        for i in range(2, N + 1):
            if is_prime(i):
                print(i)
    except ValueError:
        pass