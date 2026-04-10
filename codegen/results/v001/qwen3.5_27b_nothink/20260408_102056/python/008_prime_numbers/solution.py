import sys
import math

def sieve_of_eratosthenes(limit: int) -> None:
    if limit < 2:
        return

    # Boolean array where index represents the number
    # Initially assume all numbers are prime
    is_prime = [True] * (limit + 1)
    is_prime[0] = is_prime[1] = False

    # Optimization: iterate only up to sqrt(limit)
    for i in range(2, int(math.isqrt(limit)) + 1):
        if is_prime[i]:
            # Mark multiples of i starting from i*i
            for j in range(i * i, limit + 1, i):
                is_prime[j] = False

    # Print all numbers marked as prime
    for i in range(2, limit + 1):
        if is_prime[i]:
            print(i)

def main():
    if len(sys.argv) != 2:
        # If no argument is provided, do nothing or handle error as per spec
        # Spec says "Accept a single integer argument", implying we expect one.
        # If missing, we simply do nothing or exit.
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, do nothing or exit
        return

    sieve_of_eratosthenes(n)

if __name__ == "__main__":
    main()