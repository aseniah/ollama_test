import sys

def is_prime(n):
    """Check if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    
    # Check odd divisors up to sqrt(n)
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True

def main():
    # Get the integer argument from command line
    if len(sys.argv) != 2:
        print("Usage: python script.py N")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer")
        sys.exit(1)
    
    # Print all prime numbers up to and including n
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()