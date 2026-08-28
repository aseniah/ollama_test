import sys

def get_primes(n):
    """Returns a list of primes up to and including n using the Sieve of Eratosthenes."""
    if n < 2:
        return []
    
    # Create a boolean array "prime[0..n]" and initialize
    # all entries it as true. A value in prime[i] will
    # finally be false if i is Not a prime, else true.
    primes_mask = [True] * (n + 1)
    primes_mask[0] = primes_mask[1] = False
    
    p = 2
    while (p * p <= n):
        # If primes_mask[p] is not changed, then it is a prime
        if primes_mask[p]:
            # Updating all multiples of p
            for i in range(p * p, n + 1, p):
                primes_mask[i] = False
        p += 1
    
    # Generate list of primes from the mask
    return [p for p in range(2, n + 1) if primes_mask[p]]

def main():
    # Check if an argument is provided
    if len(sys.argv) != 2:
        return

    try:
        # Parse the input argument
        n = int(sys.argv[1])
        
        # Get the list of primes
        primes = get_primes(n)
        
        # Print each prime on a new line
        for prime in primes:
            print(prime)
            
    except ValueError:
        # If the argument is not a valid integer, exit silently or handle error
        pass

if __name__ == "__main__":
    main()