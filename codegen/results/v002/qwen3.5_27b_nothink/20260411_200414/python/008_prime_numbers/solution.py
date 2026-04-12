import sys
import math

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    limit = int(math.isqrt(n)) + 1
    for i in range(3, limit, 2):
        if n % i == 0:
            return False
    return True

def main():
    if len(sys.argv) != 2:
        return  # Exit silently if argument count is wrong

    try:
        n = int(sys.argv[1])
    except ValueError:
        return  # Exit silently if argument is not an integer

    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()