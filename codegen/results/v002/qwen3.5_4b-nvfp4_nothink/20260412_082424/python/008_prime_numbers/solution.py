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
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>", file=sys.stderr)
        sys.exit(1)

    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Invalid input. Please provide an integer.", file=sys.stderr)
        sys.exit(1)

    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)

    # If there are no primes (e.g., N < 2), print nothing as per requirements.
    # If primes exist, print each one on a new line.
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()