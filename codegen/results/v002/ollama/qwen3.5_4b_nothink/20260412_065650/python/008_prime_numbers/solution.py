import sys

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    for i in range(3, int(num ** 0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.stderr.write("Usage: python prime_numbers.py <N>\n")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.stderr.write("Error: N must be an integer.\n")
        sys.exit(1)
    
    if n < 0:
        # No primes for negative numbers
        sys.exit(0)
    
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)