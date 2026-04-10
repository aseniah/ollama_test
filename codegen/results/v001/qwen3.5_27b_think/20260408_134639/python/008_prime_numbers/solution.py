import sys
import math

def is_prime(n):
    """Check if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(math.sqrt(n)) + 1, 2):
        if n % i == 0:
            return False
    return True

def main():
    if len(sys.argv) != 2:
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(1)
    
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()