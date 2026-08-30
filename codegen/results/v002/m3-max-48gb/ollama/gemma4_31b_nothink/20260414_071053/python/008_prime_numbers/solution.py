import sys

def print_primes_up_to_n():
    """
    Reads an integer N from command line arguments and prints 
    all prime numbers up to and including N.
    """
    # Ensure a command line argument is provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the argument as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, do nothing
        return

    # Primes are greater than 1
    if n < 2:
        return

    # Implement the Sieve of Eratosthenes for efficiency
    # Create a boolean array "prime[0..n]" and initialize
    # all entries it as true. A value in prime[i] will
    # finally be false if i is Not a prime, else true.
    primes_mask = [True] * (n + 1)
    primes_mask[0] = primes_mask[1] = False

    p = 2
    while (p * p <= n):
        # If primes_mask[p] is not changed, then it is a prime
        if primes_mask[p]:
            # Update all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                primes_mask[i] = False
        p += 1

    # Print all indices that remained True
    for p in range(2, n + 1):
        if primes_mask[p]:
            print(p)

if __name__ == "__main__":
    print_primes_up_to_n()