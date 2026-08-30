import sys

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True

def main():
    # Read the integer N from command line argument
    if len(sys.argv) != 2:
        print("Error: Provide exactly one integer argument.", file=sys.stderr)
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer.", file=sys.stderr)
        sys.exit(1)
    
    # Print all prime numbers up to and including N
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()