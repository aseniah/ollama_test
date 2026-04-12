import sys

def find_primes(n):
    """
    Uses the Sieve of Eratosthenes to find all primes up to and including n.
    """
    if n < 2:
        return []

    # Create a boolean array "is_prime[0..n]" and initialize
    # all entries it as true. A value in is_prime[i] will
    # finally be false if i is Not a prime, else true.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False

    for p in range(2, int(n**0.5) + 1):
        # If is_prime[p] is not changed, then it is a prime
        if is_prime[p]:
            # Updating all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False

    # Return list of all indices that are still marked as True
    return [p for p, prime in enumerate(is_prime) if prime]

def main():
    # Check if the argument was provided
    if len(sys.argv) != 2:
        return

    try:
        # Parse the command line argument to an integer
        n = int(sys.argv[1])
        
        # Calculate primes
        primes = find_primes(n)
        
        # Print each prime on a new line
        for prime in primes:
            print(prime)
            
    except ValueError:
        # If the argument is not an integer, we exit silently 
        # or handle as per requirement.
        pass

if __name__ == "__main__":
    main()