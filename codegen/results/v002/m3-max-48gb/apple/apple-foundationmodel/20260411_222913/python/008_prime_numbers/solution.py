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

def print_primes(n):
    """Print all prime numbers up to and including n."""
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    # Check if the correct number of arguments is provided
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
    else:
        try:
            n = int(sys.argv[1])
            print_primes(n)
        except ValueError:
            print("Please provide a valid integer.")