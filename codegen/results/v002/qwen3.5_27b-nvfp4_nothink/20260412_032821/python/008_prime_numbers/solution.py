import sys
import math

def main():
    # Check if a command line argument was provided
    if len(sys.argv) != 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        # If the input is not a valid integer, do nothing or exit quietly
        return

    if n < 2:
        # No primes below 2
        return

    # Sieve of Eratosthenes to find all primes up to n
    # Initialize a boolean list where index represents the number
    is_prime = [True] * (n + 1)
    is_prime[0] = False
    is_prime[1] = False

    # Start from 2, go up to sqrt(n)
    for i in range(2, int(math.sqrt(n)) + 1):
        if is_prime[i]:
            # Mark multiples of i starting from i*i
            for j in range(i * i, n + 1, i):
                is_prime[j] = False

    # Print all prime numbers
    for num in range(2, n + 1):
        if is_prime[num]:
            print(num)

if __name__ == "__main__":
    main()