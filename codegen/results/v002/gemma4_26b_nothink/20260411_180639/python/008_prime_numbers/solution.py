import sys

def get_primes(n):
    """
    Returns a list of all prime numbers up to and including n 
    using the Sieve of Eratosthenes.
    """
    if n < 2:
        return []
    
    # Create a boolean array "is_prime[0..n]" and initialize
    # all entries it as True. A value in is_prime[i] will
    # finally be False if i is Not a prime, else True.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False
    
    p = 2
    while (p * p <= n):
        # If is_prime[p] is not changed, then it is a prime
        if is_prime[p]:
            # Update all multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
        p += 1
    
    # Collect all indices that are still marked as True
    return [p for p in range(n + 1) if is_prime[p]]

def main():
    # Check if an argument was provided
    if len(sys.argv) != 2:
        # Depending on requirements, you might want to exit or show usage
        # Here we simply exit as per the instruction to "Accept a single integer"
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
        
        # Generate the primes
        primes = get_primes(n)
        
        # Print each prime on a new line
        for prime in primes:
            print(prime)
            
    except ValueError:
        # If the argument is not an integer, do nothing
        pass

if __name__ == "__main__":
    main()