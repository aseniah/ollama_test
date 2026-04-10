import sys

def is_prime(num):
    """Check if a number is prime."""
    if num <= 1:
        return False
    if num <= 3:
        return True
    if num % 2 == 0 or num % 3 == 0:
        return False
    i = 5
    while i * i <= num:
        if num % i == 0 or num % (i + 2) == 0:
            return False
        i += 6
    return True

def print_primes_up_to(N):
    """Print prime numbers up to and including N."""
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)
    
    N = int(sys.argv[1])
    print_primes_up_to(N)