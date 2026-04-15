import sys

def print_primes_up_to_n():
    """
    Reads an integer N from command line arguments and prints all 
    prime numbers up to and including N, one per line.
    """
    # Check if a command line argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the argument as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, print nothing/exit
        return

    # Prime numbers are greater than 1
    if n < 2:
        return

    # Sieve of Eratosthenes algorithm to find all primes up to N
    # Create a boolean array "prime[0..n]" and initialize
    # all entries it as true. A value in prime[i] will
    # finally be false if i is Not a prime, else true.
    primes_mask = [True] * (n + 1)
    primes_mask[0] = primes_mask[1] = False
    
    p = 2
    while (p * p <= n):
        # If primes_mask[p] is not changed, then it is a prime
        if primes_mask[p] == True:
            # Update all multiples of p
            for i in range(p * p, n + 1, p):
                primes_mask[i] = False
        p += 1

    # Print all prime numbers
    for p in range(2, n + 1):
        if primes_mask[p]:
            print(p)

if __name__ == "__main__":
    print_primes_up_to_n()